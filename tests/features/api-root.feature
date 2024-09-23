Feature: API v1
  Scenario: Load Readme
    When I request "/"
    Then I should receive Readme document
