package uinames

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

type Name struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

type NameExtra struct {
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Gender     string     `json:"gender"`
	Region     string     `json:"region"`
	Age        int        `json:"age"`
	Title      string     `json:"title"`
	Phone      string     `json:"phone"`
	Birthday   Birthday   `json:"birthday"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	CreditCard CreditCard `json:"credit_card"`
	Photo      string     `json:"photo"`
}

type Birthday struct {
	DayMonthYear string `json:"dmy"`
	MonthDayYear string `json:"mdy"`
	Raw          int    `json:"raw"`
}

type CreditCard struct {
	Expiration string `json:"expiration"`
	Number     string `json:"number"`
	Pin        int    `json:"pin"`
	Security   int    `json:"security"`
}

type Names []Name

type NamesExtra []NameExtra

type Options struct {
	Amount int    `url:"amount,omitempty"`
	Gender string `url:"gender,omitempty"`
	Region string `url:"region,omitempty"`
	MinLen int    `url:"minlen,omitempty"`
	MaxLen int    `url:"maxlen,omitempty"`
}

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

func getBytes(o *Options, extra bool) ([]byte, error) {
	url := "https://uinames.com/api/?"
	if o != nil {
		v, err := query.Values(*o)
		if err != nil {
			return []byte{}, err
		}
		url += v.Encode()
	}

	if extra {
		url += "&ext"
	}
	return get(url)
}

func GetName(o *Options) (Name, error) {
	bytes, err := getBytes(o, false)
	if err != nil {
		return Name{}, err
	}

	name := Name{}
	err = json.Unmarshal(bytes, &name)
	return name, nil
}

func GetNameExtra(o *Options) (NameExtra, error) {
	bytes, err := getBytes(o, true)
	if err != nil {
		return NameExtra{}, err
	}

	name := NameExtra{}
	err = json.Unmarshal(bytes, &name)
	return name, nil
}

func GetNames(o *Options) (Names, error) {
	bytes, err := getBytes(o, false)
	if err != nil {
		return Names{}, err
	}

	names := Names{}
	if o != nil && o.Amount > 1 {
		err = json.Unmarshal(bytes, &names)
		if err != nil {
			return Names{}, err
		}
	} else {
		name := Name{}
		err = json.Unmarshal(bytes, &name)
		if err != nil {
			return Names{}, err
		}
		names = Names{name}
	}
	return names, nil
}

func GetNamesExtra(o *Options) (NamesExtra, error) {
	bytes, err := getBytes(o, true)
	if err != nil {
		return NamesExtra{}, err
	}

	names := NamesExtra{}
	if o != nil && o.Amount > 1 {
		err = json.Unmarshal(bytes, &names)
		if err != nil {
			return NamesExtra{}, err
		}
	} else {
		name := NameExtra{}
		err = json.Unmarshal(bytes, &name)
		if err != nil {
			return NamesExtra{}, err
		}
		names = NamesExtra{name}
	}
	return names, nil
}
