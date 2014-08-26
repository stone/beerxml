package beerxml

import (
	"testing"
)

func TestUnmarshalBeerXml(t *testing.T) {
	t.Log("Unmarshal beerxml recipes.xml")
	_, err := NewBeerXmlFromFile("testfiles/recipes.xml")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestUnmarshalBeerSmithExport(t *testing.T) {
	t.Log("Unmarshal beerxml drsmurto.xml")
	_, err := NewBeerXmlFromFile("testfiles/drsmurto.xml")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestRecipesRecipe(t *testing.T) {
	//t.Log("Unmarshal test ... ")
	bxml, err := NewBeerXmlFromFile("testfiles/recipes.xml")
	if err != nil {
		t.Error(err)
		return
	}

	if bxml.Recipes[0].Name != "Burton Ale" {
		t.Error("Recipe.Name")
	}

	if bxml.Recipes[0].Hops[0].Hsi != "35.0" {
		t.Error("Recipe.Hops.Hsi")
	}

	if bxml.Recipes[0].Fermentables[0].Amount != 3.628736 {
		t.Error("Recipe.Fermentables.Amount")
	}

	if bxml.Recipes[0].Mash.MashSteps[0].StepTime != 45 {
		t.Error("Recipes.Mash.MashSteps.StepTime")
	}
}
