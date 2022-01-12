# JSON parser using GO

## An example of parsing json data in go, when you already know the schema of the data 


### **Example input**:

```json
{
    "num_listings":"36",
    "City":"Toronto",
    "Neighbourhood":"Downtown",
    "Area":"Moss Park",
    "Listings":[    
        {
            "MLS_num":"C5467882",
            "Price":"3,200",
            "Address":"1007 - 219 Dundas  St E",
            "Bedrooms":"3BD",
            "Bathrooms":"2BA",
            "Parking":"0 Parking",
            "Size":"800-899 sqft",
            "timeListed":"11 hours",
            "Brokerage":"LIVING REALTY INC., BROKERAGE",
            "timeScraped":"01/10/2022-00:13:38"
         }
    ]     
}

```
### **Example output**:
```json
Listings :  [map[Address:1007 - 219 Dundas  St E Bathrooms:2BA Bedrooms:3BD Brokerage:LIVING REALTY INC., BROKERAGE MLS_num:C5467882 Parking:0 Parking Price:3,200 Size:800-899 sqft timeListed:11 hours timeScraped:01/10/2022-00:13:38]]
```
