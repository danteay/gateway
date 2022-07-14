## v2.1.2 (2022-07-14)


- fix: tests
- fix: response object should be an api gateway response

## v2.1.1 (2022-07-14)


- fix: event parse from api gatewahandler

## v2.1.0 (2022-07-14)


- fix: release
- feat: add handler options to decorate lambda handler execution

## v1.1.2 (2020-10-20)


- Merge pull request #37 from wolfeidau/fix_request_uri
- fix requestURI for v2 gateway events
- fix requestURI for v2 gateway events
- As per #34
- Merge pull request #34 from jainpiyush19/master
- Fix Request.RequestURI missing query params
- Merge pull request #36 from nobu-k/add-cookies
- Add Cookies to Headers in v2
- Add Cookies to Headers in v2
- Fix Request.RequestURI missing query params
- Cf. https://github.com/apex/gateway/issues/26#issuecomment-689386480
- Merge pull request #32 from wolfeidau/feat_github_actions
- feat(build) add github actions and updated README with options
- feat(build) add github actions and updated README with options
- Merge pull request #30 from wolfeidau/feat_extend_apigwv2
- Add support for 2.0 apigw request/response
- feat(apigw) Add support for 2.0 request/response
- * This adds a new v2 module in a subfolder
* Ensure tests use canonical format used in 1.14
* Upgrade dependencies
* Use the Handler interface to add some extensibility
- Add application/javascript as text mime type (#28)
- Set Request.RequestURI (#27)
- Fixes #26
- Remove unnecessary import in README (#25)
- add go.mod. Closes #21
- add support for multi param and multi header (#20)
- * add support for multi param and multi header

adds support for multiple parameter for requests
adds support for multiple headers in request and response
adds tests

* simplify assigning query parameters

just set query[key] to value instead of calling Add

## v1.1.1 (2018-08-17)


- Release v1.1.1
- Merge pull request #13 from dstrelau/passthrough-context
- Pass-through incoming ctx to Request
- pass-through incoming ctx to Request
- At this point, the request has just been created and `req.Context()` is
just the `Background()` context. Pass through the `ctx` that we got as a
parameter so that any values / deadlines / etc are carried through.

## v1.1.0 (2018-07-05)


- Merge pull request #11 from alikor/patch-1
- Remove brackets on license to support linting tools
- Update LICENSE
- Merge pull request #2 from olivoil/feature/context
- add support for `APIGatewayProxyRequestContext`
- add ci.yml
- Merge pull request #7 from wolfeidau/feat_add_closenotify_support
- feat(Response) Add CloseNotify support.
- feat(Response) Add CloseNotify support.
- This is required to use `httputils.ReverseProxy` and enable proxying of requests otherwise it returns the following error:
- > *gateway.ResponseWriter is not http.CloseNotifier: missing method CloseNotify
- It is used with `context.Context` to interupt in flight requests when the response is ended.
- add license
- add Example()
- add response tests
- add binary body test
- add request tests

## v1.0.0 (2018-01-25)


- Merge pull request #6 from wolfeidau/fix_text_mimetypes
- fix(Response) Fix for autodetect text mime types.
- fix(Response) Fix for autodetect text mime types.
- * Changed matching to allow for trailing encoding.
* Added a test case for the updated function.
- Merge branch 'master' into feature/context
- add return to readme example
- Merge pull request #3 from wolfeidau/feature_xray_header
- feat(xray) Pass through the xray trace identifier.
- feat(xray) Pass through the xray trace identifier.
- As per the suggestion by @bmoffatt https://github.com/aws/aws-lambda-go/issues/15#issuecomment-358879698
- fix readme
- add api gateway request context to http request context

## v0.0.1 (2018-01-17)


- Initial commit
