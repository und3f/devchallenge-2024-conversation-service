Feature: Category
  Scenario: Create category
    When I make a request to create sample category
    Then I should receive category created success response

  Scenario: List categories
    When I request "/category"
    Then I should receive list of all conversation topics
    And  I should see default conversation topics

  Scenario: Update category
    When I request to update previously created sample category
    Then category should be updated

  Scenario: Delete category
    When I request to delete previously created sample category
    Then category should be unavailable
