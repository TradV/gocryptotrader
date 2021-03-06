package currency

import (
	"encoding/json"
	"testing"
)

func TestCurrenciesUnmarshalJSON(t *testing.T) {
	var unmarshalHere Currencies
	expected := "btc,usd,ltc,bro,things"
	encoded, err := json.Marshal(expected)
	if err != nil {
		t.Fatal("Currencies UnmarshalJSON() error", err)
	}

	err = json.Unmarshal(encoded, &unmarshalHere)
	if err != nil {
		t.Fatal("Currencies UnmarshalJSON() error", err)
	}

	err = json.Unmarshal(encoded, &unmarshalHere)
	if err != nil {
		t.Fatal("Currencies UnmarshalJSON() error", err)
	}

	if unmarshalHere.Join() != expected {
		t.Errorf("Currencies UnmarshalJSON() error expected %s but received %s",
			expected, unmarshalHere.Join())
	}
}

func TestCurrenciesMarshalJSON(t *testing.T) {
	quickStruct := struct {
		C Currencies `json:"amazingCurrencies"`
	}{
		C: NewCurrenciesFromStringArray([]string{"btc", "usd", "ltc", "bro", "things"}),
	}

	encoded, err := json.Marshal(quickStruct)
	if err != nil {
		t.Fatal("Currencies MarshalJSON() error", err)
	}

	expected := `{"amazingCurrencies":"btc,usd,ltc,bro,things"}`
	if string(encoded) != expected {
		t.Errorf("Currencies MarshalJSON() error expected %s but received %s",
			expected, string(encoded))
	}
}
