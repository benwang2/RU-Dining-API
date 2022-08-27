# Rutgers Dining API
[![Run on Repl.it](https://replit.com/badge/github/benwang2/RU-Dining-API)](https://replit.com/new/github/benwang2/RU-Dining-API)

An unofficial Rutgers Dining web API that web scrapes and transforms the Rutgers Dining Services Menu to an easily consumable JSON format.

## About the Project
This web API was hastily created to supplement an Android application that shows the current Rutgers dining hall menus. As of August 27, 2022, this Android application has not been created.

## Usage
This project can be deployed on [Replit.com](https://replit.com/) for demonstration purposes.

By default, the API endpoint can be found at `localhost:3333/api/menu`.

An example URL query and response would look like the following:

`http://localhost:3333/api/menu?location=busch&meal=lunch&date=09/02/2022`

<sub>Note that the example response below is significantly smaller than an actual response.</sub>
```json
{
  "StatusCode": 200,
  "Status": "OK",
  "URL": "http://menuportal.dining.rutgers.edu/FoodPro/pickmenu.asp?locationNum=04\u0026mealName=Lunch\u0026dtdate=09/02/2022",
  "Menu": [
    {
      "Name": "BREAKFAST MEATS",
      "Items": [
        {
          "Name": "PORK SAUSAGE LINKS",
          "Info": "label.asp?RecNumAndPort=195018%2A3"
        },
        {
          "Name": "TURKEY BACON",
          "Info": "label.asp?RecNumAndPort=705281%2A4"
        },
        {
          "Name": "VEGETARIAN BREAKFAST PATTIES",
          "Info": "label.asp?RecNumAndPort=143000%2A2"
        }
      ]
    },
        ...
    {
      "Name": "LUNCH TO GO",
      "Items": [
        {
          "Name": "CHEESE STUFFED RIGATONI",
          "Info": "label.asp?RecNumAndPort=140141%2A6"
        },
        {
          "Name": "CHICKEN CAESAR WRAP",
          "Info": "label.asp?RecNumAndPort=600929%2A1"
        },
        {
          "Name": "CHIPOTLE BLACK BEAN BURGER",
          "Info": "label.asp?RecNumAndPort=019105%2A1"
        },
      ]
    }
  ]
}
```