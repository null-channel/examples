# News feed 

## What's this?

An API that returns authentic news :wink: 


## Setup 
* clone this repo and cd into the news-feed directory 

* run the api `go run main.go --port <port-number>`


## Endpoints


### Endpoint: /api/news/feed 

```bash
curl localhost:8080/api/news/feed
```

### Response

```json
[
 {
    "title": "information reveal politely sithence east",
    "author": "Augusta Ortiz",
    "date": "7-May-1925",
    "description": "set radio,web-enabled kinfolk prefer.",
    "body": "Vinegar brooklyn ennui cardigan master. Five dollar toast asymmetrical portland authentic shabby chic. Portland wayfarers crucifix seitan bespoke. Beard drinking pop-up locavore normcore. Blog food truck farm-to-table tattooed normcore.,Vice irony artisan hella vinyl. Chartreuse deep v seitan kinfolk taxidermy. Tote bag normcore narwhal sriracha typewriter. Drinking neutra austin deep v seitan. Tacos mustache 3 wolf moon yr authentic.,Etsy twee cliche chartreuse tofu. Franzen marfa synth iPhone lumbersexual. Gentrify ennui pug 8-bit sartorial. Fingerstache salvia meh ennui franzen. Echo cliche vinyl vice offal.,Echo 90's YOLO slow-carb meditation. Bitters fashion axe intelligentsia ethical forage. Leggings stumptown humblebrag health taxidermy. Trust fund +1 deep v raw denim ennui. Roof franzen kogi typewriter seitan.,Kogi stumptown deep v hashtag butcher. Taxidermy vinyl slow-carb post-ironic Godard. Plaid mumblecore tilde hashtag deep v. Heirloom plaid Thundercats polaroid pabst. Meh put a bird on it cred cliche chillwave."
  },
  {
    "title": "hand match logically aprés exhibition",
    "author": "Dallas Sauer",
    "date": "29-February-1971",
    "description": "aim reduce,multi-state celiac total.",
    "body": "Cleanse narwhal flexitarian retro scenester. Flannel flexitarian cornhole tilde kombucha. Vegan five dollar toast health truffaut +1. Direct trade pickled pickled pour-over asymmetrical. Taxidermy street brunch shoreditch ennui.,Taxidermy VHS gastropub DIY kinfolk. Squid beard occupy leggings scenester. Literally kale chips irony umami actually. Trust fund sriracha park try-hard chambray. Migas kogi butcher cred chia.,Synth wolf tote bag pug hoodie. Brunch deep v pour-over post-ironic fixie. Meggings kale chips cronut Wes Anderson quinoa. Literally literally offal locavore tote bag. Deep v photo booth fingerstache sartorial church-key.,Aesthetic tilde meh venmo polaroid. Normcore paleo tote bag forage wolf. Vinegar hella chia echo shabby chic. Migas gastropub umami put a bird on it locavore. Stumptown tilde jean shorts celiac green juice.,Mlkshk migas offal synth 90's. Flexitarian bicycle rights next level iPhone semiotics. Poutine lo-fi put a bird on it marfa hammock. Salvia Godard forage marfa drinking. Blog taxidermy single-origin coffee mixtape stumptown."
  },
  {
    "title": "independence have impatiently neath cash",
    "author": "Jadon Bauch",
    "date": "31-March-1940",
    "description": "choose extent,bandwidth-monitored gentrify organise.",
    "body": "Vice brooklyn ethical waistcoat flexitarian. Heirloom kombucha seitan kinfolk salvia. Banjo locavore carry messenger bag chambray. You probably haven't heard of them seitan portland pork belly beard. Shoreditch viral tumblr neutra etsy.,Everyday bespoke flexitarian messenger bag Yuccie. Banh mi fixie hammock hella PBR&B. Lumbersexual crucifix food truck plaid authentic. You probably haven't heard of them hammock marfa lo-fi Godard. Mixtape microdosing meditation Godard goth.,Cornhole post-ironic fashion axe tote bag leggings. Bicycle rights cornhole wolf photo booth cold-pressed. +1 lomo freegan PBR&B vegan. Marfa listicle butcher +1 letterpress. Gluten-free freegan scenester semiotics asymmetrical.,Tilde ennui taxidermy migas tattooed. Semiotics PBR&B cleanse Thundercats kogi. Swag keffiyeh typewriter authentic art party. Craft beer bicycle rights portland sriracha quinoa. Taxidermy readymade Yuccie put a bird on it vinegar.,Sriracha schlitz flannel cray gentrify. Pabst mustache swag cleanse gluten-free. Fingerstache humblebrag tofu forage sustainable. Typewriter poutine park hammock seitan. Normcore seitan aesthetic readymade ugh."
  },

]


```

### Endpoint: /api/news/breif
for a simple news digest

```bash
curl localhost:8080/api/news/brief
```

### Response

```json
{
  "title": "japan account rarely amid belief",
  "description": "agree thank,contingency farm-to-table group."
}

```
