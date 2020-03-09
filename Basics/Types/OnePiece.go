package main

func main() {
	crew := mugiwara{"Luffy","Zoro","Nami","Sanji","Usoup"}
	crew = append(crew, "Robin")
	crew = append(crew, "Frankie")
	crew.print()
}
