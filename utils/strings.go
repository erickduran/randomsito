package utils

var langStrings = map[string]map[string]string {
	"greet": {
		"en": "randomsito says hi 👋",
		"es": "randomsito dice hola 👋",
	},
	"options": {
		"en": "options",
		"es": "opciones",
	},
	"choose": {
		"en": "choose",
		"es": "dale",
	},
	"add": {
		"en": "add student",
		"es": "agregar alumnito",
	},
	"askName": {
		"en": "what's your name?",
		"es": "¿cómo te llamas?",
	},
	"added": {
		"en": "added!",
		"es": "listo",
	},
	"points": {
		"en": "points",
		"es": "puntitos",
	},
	"exit": {
		"en": "bye",
		"es": "adiós",
	},
	"bye": {
		"en": "bye!",
		"es": "bai",
	},
	"connected":{
		"en": "connected to MongoDB",
		"es": "conectado a MongoDB",
	},
	"connecting":{
		"en": "connecting to MongoDB",
		"es": "conectando a MongoDB",
	},
	"refresh": {
		"en": "refreshing...",
		"es": "reiniciando...",
	},
	"selected": {
		"en": "randomsito chose",
		"es": "randomsito elegió a",
	},
	"selectClassroom": {
		"en": "choose a classroom",
		"es": "escoge un grupo",
	},
	"addClassroom": {
		"en": "create a classroom",
		"es": "agregar grupo",
	},
	"askClassroomName": {
		"en": "classroom name:",
		"es": "nombre del grupo:",
	},
	"noStudents": {
		"en": "there's no students in this classroom",
		"es": "no hay alumnos en este grupo",
	},
}

func GetString(language string, indentifier string) string {
	return langStrings[indentifier][language]
}
