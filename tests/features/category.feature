Feature: Category
  Scenario: Create category
    When I make a request to create sample category
    Then I should receive category created success response

  Scenario: List categories
    When I request "/category"
    Then I should receive list of all conversation topics
