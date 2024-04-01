⬆️ For table of contents, click the above icon

xs.fi ("accessify"): semantic, persistent URL namespace and a metadata specification.

Let's make URLs and metadata beautiful again.

We offer URLs like [https://xs.fi/movie/by-imdb/tt12593682/imdb](https://xs.fi/movie/by-imdb/tt12593682/imdb)

We also plan to be a metadata specification to provide guide how to record metadata for your precious

If you recognize yourself as a data hoarder, you'll feel right at home.


URL model
---------

| Path after `https://xs.fi` | Example                     | Components      | Description |
|-----------------------|-----------------------------|-----------------|-------------|
| `/<namespace>` | `https://xs.fi/movie` | `namespace=movie` | We are going to be referencing a movie |
| `/<namespace>/<query>` | `https://xs.fi/movie/by-imdb` | `namespace=movie`, `query=by-imdb` | Movie by IMDb ID |
| `/<namespace>/<query>/<identifier>` | `https://xs.fi/movie/by-imdb/tt12593682` | `namespace=movie`, `query=by-imdb`, `id=tt12593682` | Movie by IMDb ID tt12593682 |
| `/<namespace>/<query>/<identifier>/<resource>` | `https://xs.fi/movie/by-imdb/tt12593682/imdb` | `namespace=movie`, `query=by-imdb`, `id=tt12593682`, `resource=imdb` | IMDb page for movie by IMDb ID tt12593682 |

In pseudo you can think of our backend data model like this:

```json
{
	"namespaces": {
		"movie": {
			"queries": [
				{
					"query": "by-imdb"
				}
			],
			"resources": [
				{
					"resource": "imdb",
					"description": "IMDb page",
					"redirect_to": "https://www.imdb.com/title/<IMDb ID>"
				}
			]
		},
		"tv-series": {
			"queries": [
				{
					"query": "by-imdb"
				}
			],
			"resources": [
				{
					"resource": "imdb",
					"description": "IMDb page",
					"redirect_to": "https://www.imdb.com/title/<IMDb ID>"
				}
			]
		}
	}
}
```

- `namespace` is the first part after xs.fi: `https://xs.fi/<namespace>`. Example namespace is `movie`.
- Namespace is *usually* followed by `query` so the format usually is `https://xs.fi/<namespace>/<query>` After the namespace usually comes `by-<identifier type>` query.
  Example is `https://xs.fi/<namespace>/by-imdb` which means we are referencing a node based on its IMDB ID.
  For `by-imdb` namespaces which make sense are `movie` or `tv-series` so namespace+query combos read like `movie/by-imdb` or `tv-series/by-imdb`
- *Identifier* 
- *Resource*


Why?
----

Practice has shown that not all people think URLs are beautiful or important: they'll change their web application UI and break old URLs.

So if IMDb ever changes the title URL to something else than `https://www.imdb.com/title/<ID>` then
xs.fi can change that in the code to link to the new URL


Audience
--------

### Human consumption

Currently the URLs are meant for **human consumption**, i.e. it is expected the URLs to be opened in a browser.


### API access

> [!WARNING]  
> Not implemented. This *might* be added later.


That will be handled by [Content negotiation](https://en.wikipedia.org/wiki/Content_negotiation),
concretely something like `Accept: application/json` to distinguish opt-in to programmatic access instead of human consumption.


TODO
----

- Support "xs.fi:ification" of URLs: give `https://www.imdb.com/title/tt12593682` as input and it
  should spit you back `https://xs.fi/movie/by-imdb/tt12593682/imdb` and also related resources like TMDB.


Promises
--------

xs.fi will always:

- Stay
	* our largest value proposition is to not break your links and not annoy you
- Be free
- Not show advertisements
	* or if we ever need to resort to ads then it will be based on
	  [Acceptable Ads](https://acceptableads.com/standard/), preferably the Text ads variant.
