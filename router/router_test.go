package router

import (
  "testing"
  "github.com/a8m/expect"
)

func TestStringPrefix (t* testing.T) {
  expect := expect.New(t)
  expect(stripPrefix("", "")).To.Equal("");
  expect(stripPrefix("/", "")).To.Equal("/");
  expect(stripPrefix("/a", "/a")).To.Equal("/");
  expect(stripPrefix("/a/", "/a")).To.Equal("/");
  expect(stripPrefix("/a/b", "/a")).To.Equal("/b");
  expect(stripPrefix("/a/", "/a/")).To.Equal("/");
  expect(stripPrefix("/a/b", "/a/")).To.Equal("/b");
}

func TestRouteBasic (t* testing.T) {
  expect := expect.New(t)
  router := Create(&RouterOptions{
    CaseType: "camel",
    Method: "GET",
    Sensitive: true,
  })
  jsonText := `{
    "modals": {
      "Home": {
        "routes": {
          "default": {},
          "relative": {
            "path": "relativeRoute"
          },
          "absolute": {
            "path": "/absoluteRoute"
          },
          "user": {
            "path": "user/:id"
          }
        }
      },
      "Article": {
        "path": "art",
        "routes": [
          {"name": "list"},
          {"name": "cat", "path": "/article/cat/:id"},
          {"name": "item", "path":"item/:id"},
          {"name": "index", "path":""}
        ]
      }
    }
  }`
  router.Load(jsonText)

  expect(router.Match("/home/default", GET).Path).To.Equal("/home/default")
  expect(router.Match("/home/default/", GET).Path).To.Equal("/home/default")
  expect(router.Match("/home/relativeRoute", GET).Path).To.Equal("/home/relativeRoute")
  expect(router.Match("/home/relativeRoute/", GET).Path).To.Equal("/home/relativeRoute")
  expect(router.Match("/absoluteRoute", GET).Path).To.Equal("/absoluteRoute")
  expect(router.Match("/absoluteRoute/", GET).Path).To.Equal("/absoluteRoute")
  expect(router.Match("/home/user/1", GET).Path).To.Equal("/home/user/1")
  expect(router.Match("/home/user/1/", GET).Path).To.Equal("/home/user/1")
  expect(router.Match("/art", GET).Path).To.Equal("/art")
  expect(router.Match("/art/", GET).Path).To.Equal("/art")
}

func TestRouteOptions (t* testing.T) {
  expect := expect.New(t)
  router := Create(&RouterOptions{
    CaseType: "hyphen",
    Method: "GET",
    Sensitive: false,
  })
  jsonText := `{
    "modals": {
      "UserProfile": {}
    },
    "actions": [
      {"modal":"UserProfile", "name": "HomeInfo"},
      {"modal":"UserProfile", "name": "UserAddress", "path": "UserAddress"}
    ]
  }`
  router.Load(jsonText)
  matched := router.MatchStringMethod("/user-profile/home-info", "GET")
  expect(matched.Route).To.Be.Ok()

  matched = router.MatchStringMethod("/user-PROFILE/HOME-info", "GET")
  expect(matched.Route).To.Be.Ok()

  matched = router.MatchStringMethod("/user-profile/HomeInfo", "GET")
  expect(nil == matched).To.Be.Ok()

  matched = router.MatchStringMethod("/UserProfile/home-info", "GET")
  expect(nil == matched).To.Be.Ok()

  matched = router.MatchStringMethod("/UserProfile/HomeInfo", "GET")
  expect(nil == matched).To.Be.Ok()
  
  matched = router.MatchStringMethod("/user-profile/UserAddress", "GET")
  expect(matched.Route).To.Be.Ok()

}

func TestRouter(t * testing.T) {
  expect := expect.New(t)
  jsonText :=  `{
    "modals": [
      {
        "name": "UserProfile",
        "id": 1,
        "routes": [
          {
            "name": "UserAddress"
          }
        ]
      }
    ],
    "actions": [
      {
        "modal": "1",
        "name": "HomeInfo"
      }
    ]
  }`
  var router = Create(&RouterOptions{
    CaseType: "hyphen",
    Method: "GET",
    Sensitive: false,
  })
  router.Load(jsonText)

  matched := router.MatchStringMethod("/user-profile/home-info", "GET")
  expect(matched.Route).To.Be.Ok()

  matched = router.MatchStringMethod("/user-PROFILE/HOME-info", "GET")
  expect(matched.Route).To.Be.Ok()

  matched = router.MatchStringMethod("/user-profile/HomeInfo", "GET")
  expect(nil == matched).To.Be.Ok()

  matched = router.MatchStringMethod("/UserProfile/home-info", "GET")
  expect(nil == matched).To.Be.Ok()

  matched = router.MatchStringMethod("/UserProfile/HomeInfo", "GET")
  expect(nil == matched).To.Be.Ok()
  
  matched = router.MatchStringMethod("/user-profile/user-address", "NONE")
  expect(nil == matched).To.Be.Ok()

  matched = router.MatchStringMethod("/user-profile/user-address", "OPTIONS")
  expect(matched.Route).To.Be.Ok()

}
