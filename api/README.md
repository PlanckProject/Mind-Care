# This directory contains the Golang code of the RESTful API server of [Mind Care](https://mindcare.page).

##### TODO List:
- Build complete set of CRUD routes for service provider data.
- Secure the mutation API calls of service provider data. (User registration? Basic credentials hardcoded or from config?)
- Do not expose Mongo "_id" as ID of the document, come up with a legible, rememberable, and secure approach. Potentially form ID from the name, etc.
- Reduce dependency on ENV variables. We have a proper config in place, why not utilize it?
- For faaaaaar future, do we need caching, shard, etc. etc. etc.? GeoJSON data is a potential candidate.