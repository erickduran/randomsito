package cmd

import (
	"os"
	"fmt"
	"time"
	"log"
	"bufio"
	"strings"
	"context"
	"math/rand"
	"github.com/spf13/cobra"
	"github.com/manifoldco/promptui"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/erickduran/randomsito/utils"
)

var classroom *mongo.Collection
var ctx context.Context

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Normal interactive session",
	Long: `This command starts a normal interactive session
	where you can randomly select students in a
	round-robin fashion.`,
	Run: func(cmd *cobra.Command, args []string) {	
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		fmt.Println("⏳", utils.GetString(Language, "connecting"))
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(ConnectionString))
		if err != nil {
		    log.Fatal(err)
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
		    log.Fatal(err)
		}

		fmt.Println("✅", utils.GetString(Language, "connected"))

		classrooms, err := client.Database(MongoDatabase).ListCollectionNames(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}

		prettyAdd := fmt.Sprintf("%s%s%s", BLUE, utils.GetString(Language, "addClassroom"), NC)
		classrooms = append(classrooms, prettyAdd)

		prompt := promptui.Select{
			Label: utils.GetString(Language, "selectClassroom"),
			Items: classrooms,
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if result == prettyAdd {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(utils.GetString(Language, "askClassroomName") + " ")
			name, _ := reader.ReadString('\n')
			result = strings.Replace(name, "\n", "", -1)
		}

		classroom = client.Database(MongoDatabase).Collection(result)
		if err != nil {
			log.Fatal(err)
		}

		for {
			prompt := promptui.Select{
				Label: utils.GetString(Language, "options"),
				Items: []string {
					utils.GetString(Language, "choose"), 
					utils.GetString(Language, "add"),
					utils.GetString(Language, "exit"),
				},
			}

			_, result, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			switch result {
				case utils.GetString(Language, "choose"):
					chosen := chooseStudent()
					if chosen != nil {
						studentPretty := fmt.Sprintf("%s%s %s%s", BLUE, chosen["name"].(string), chosen["emoji"].(string), NC)
						fmt.Println(utils.GetString(Language, "selected"), "👉", studentPretty)
					} else {
						fmt.Println("😢", utils.GetString(Language, "noStudents"))
					}
				case utils.GetString(Language, "add"):
					reader := bufio.NewReader(os.Stdin)
					fmt.Print(utils.GetString(Language, "askName") + " ")
					name, _ := reader.ReadString('\n')
					name = strings.Replace(name, "\n", "", -1)
					
					fmt.Print("emoji: ")
					emoji, _ := reader.ReadString('\n')
					emoji = strings.Replace(emoji, "\n", "", -1)

					if emoji == "" {
						emoji = "😬"
					}
					addStudent(name, emoji)
					
			    case utils.GetString(Language, "exit"):
			        fmt.Println("👋", utils.GetString(Language, "bye"))
			        os.Exit(0)
			}
		}
	},
}

func chooseStudent() bson.M {
	filter := bson.M {
		"robin": bson.M {
			"$eq": false,
		},
	}
			
	cursor, err := classroom.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	var students []bson.M
	if err = cursor.All(ctx, &students); err != nil {
	    log.Fatal(err)
	}

	if len(students) < 1 {
		cursor, err := classroom.Find(ctx, bson.M{})
		if err != nil {
		    log.Fatal(err)
		}

		if err = cursor.All(ctx, &students); err != nil {
		    log.Fatal(err)
		}

		if len(students) < 1 {
			return nil
		}

		update := bson.M {
        	"$set": bson.M{
        		"robin": false,
        	},
		}

		fmt.Println("🔄", utils.GetString(Language, "refresh"))
		_, err = classroom.UpdateMany(ctx, bson.M{}, update)
		if err != nil {
		    log.Fatal(err)
		}
	}

	rand.Seed(time.Now().Unix())
	selected := rand.Int() % len(students)

    filter = bson.M {
    	"_id": bson.M {
    		"$eq": students[selected]["_id"],
    	},
    }

    update := bson.M {
    	"$set": bson.M{
    		"robin": true,
    	},
	}

    _, err = classroom.UpdateOne(ctx, filter, update)
	if err != nil {
	    log.Fatal(err)
	}
	return students[selected]
}

func addStudent(name string, emoji string) {
	_, err := classroom.InsertOne(ctx, bson.M{
		"name": name,
		"emoji": emoji,
		"robin": false,
	})
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println("✅", utils.GetString(Language, "added"))
}

func init() {
	rootCmd.AddCommand(startCmd)
}
