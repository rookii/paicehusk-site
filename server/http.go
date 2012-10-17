package server

import (
	"encoding/json"
	"github.com/Rookii/paicehusk"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

var Templates *template.Template
var wordreg = regexp.MustCompile("[a-zA-Z]+")

func init() {

	t, err := template.New("foo").Parse(homePage)
	if err != nil {
		panic(err)
	}
	Templates = t
	router := mux.NewRouter()

	// Root handler
	router.HandleFunc("/", home)
	router.HandleFunc("/stem", stem)
	http.Handle("/", router)
}

func home(w http.ResponseWriter, r *http.Request) {
	Templates.ExecuteTemplate(w, "T", nil)
}

type Text struct {
	Text string
}

func stem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.ContentLength > 2500 {
		http.Error(w, "Content is to large", http.StatusInternalServerError)
		return
	}
	out := make(map[string]string)
	decoder := json.NewDecoder(r.Body)

	text := new(Text)

	if err := decoder.Decode(text); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	words := wordreg.FindAllString(text.Text, -1)
	for _, word := range words {
		out[strings.ToLower(word)] = paicehusk.DefaultRules.Stem(word)
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(out); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var homePage = `
{{define "T"}}

<!DOCTYPE html>
<html lang="en">
  <head>
	<meta charset="utf-8">
	<title>Paice/Husk Stemmer</title>
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta name="description" content="An implementation of the Paice/Husk Stemmer">
	<meta name="author" content="Aaron Groves">
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js"></script>
	<link href="/assets/css/bootstrap.css" rel="stylesheet">
	<style type="text/css">
	  body {
		padding-top: 60px;
		padding-bottom: 40px;
	  }
	</style>
	<link href="/assets/css/bootstrap-responsive.css" rel="stylesheet">
	<script>
		$(document).ready(function() {
			$('#stemButton').click(function() {
				stem();
			  })
		  })

		function stem(){
			var Text = $("#input").val();
			var input = {"Text": Text}
			$.ajax({
				type: "POST",
				url: "/stem",
				data: JSON.stringify(input),
			}).done(function(data) {
				var res = "| ";
				var raw = "\n\nRaw output:\n";
				$.each(data, function(key, value){
					res += key + ": " + value + " | ";
					raw += value + "\n";
				});
				$("#output").val(res + raw);
			});
		};
	</script>
  </head>

  <body>

	<div class="navbar navbar-inverse navbar-fixed-top">
	  <div class="navbar-inner">
		<div class="container">
		  <a class="btn btn-navbar" data-toggle="collapse" data-target=".nav-collapse">
			<span class="icon-bar"></span>
			<span class="icon-bar"></span>
			<span class="icon-bar"></span>
		  </a>
		  <a class="brand">Paice/Husk Stemmer</a>
		  <div class="nav-collapse collapse">
            <ul class="nav">
              <li><a href="https://github.com/Rookii/paicehusk">Source</a></li>
            </ul>
          </div><!--/.nav-collapse -->
		</div>
	  </div>
	</div>

	<div class="container">
	  <!-- Example row of columns -->
	  <div class="row">
		<div class="span6">
		  <h2>Text to Stem</h2>
		  <form>
		  <textarea class="span6" maxlength="2000" wrap="wrap" rows="25" id="input" placeholder="Place text to stem here!"></textarea>
		  </form>
		  <button id="stemButton" class="btn">Submit</button>
		</div>
		<div class="span6">
		  <h2>Stemmed Output</h2>
		  <textarea class="span6"  rows="25" readonly="readonly" id="output" placeholder="Use input Box"></textarea>
	   </div>
	  </div>

	  <hr>

	  <footer>
		<p>&copy; Aaron Groves 2012</p>
	  </footer>

	</div> <!-- /container -->
  </body>
</html>

{{end}}
`
