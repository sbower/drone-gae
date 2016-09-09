# drone-gae

Manage deployments on Google App Engine via drone.

This is currently very new and unstable.  
Please don't use it for anything important!


## Example

	build:
	  image: golang:1.7
	  commands:
	    - goapp get -t
	    - goapp test -v -cover
	  when:
	    event:
	      - push
	      - pull_request

	deploy:
	  gae:
        action: update
        project: my-gae-project
	    version: "$$COMMIT"
	    token: >
	      $$GOOGLE_CREDENTIALS

	    when:
	      event: push
	      branch: master

	  gae:
        action: set_default_version
        project: my-gae-project
	    version: "$$COMMIT"
	    token: >
	      $$GOOGLE_CREDENTIALS

	    when:
	      event: push
	      branch: master


## License

MIT.