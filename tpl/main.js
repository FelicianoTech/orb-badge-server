{{ define "js" }}

$(document).ready(function() {

	$( "form" ).submit(function(e){

		e.preventDefault();
		
		namespace = $( "input[name='namespace']" ).val();
		name = $( "input[name='name']" ).val();

		markdown = "[![CircleCI Orb Version](https://img.shields.io/badge/endpoint.svg?url=https://badges.circleci.io/orb/" + namespace + "/" + name + ")](https://circleci.com/orbs/registry/orb/" + namespace + "/" + name + ")";

		$( "textarea" ).val( markdown );
	});
});

{{ end }}
