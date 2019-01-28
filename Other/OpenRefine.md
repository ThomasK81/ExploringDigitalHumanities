Install OpenRefine from [here](http://openrefine.org/download.html).

#### Data 

http://api.digitalnz.org/v3/records.json?api_key=[your_key]&text=cartoon+rugby&and[collection][]=TAPUHI&and[year][]=2011&fields=id,title,collection&per_page=10&page=1  

#### Some Handy GREL

- value
- rowIndex
- "http://api.digitalnz.org/v3/records/" + value + ".json?api_key=[your_key]&fields=description"  
- value.parseJson().get("record").get("description")
- value.parseHtml().htmlText()
- value.replace('{"record":{"description":"', "")  
- value.partition(/[0-9]{1,2}\s\S+\s\d{4}/)[1].replace(",", "")  
- if(value.toLowercase().contains("christchurch"), 1, 0)

