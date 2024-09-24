Feature: Call
  Scenario: Create call
    When I make a request to create sample call
    Then I should receive call created success response

  Scenario: Get call
    When I make a request to get created sample call
    Then I wait for call process
    And  I should receive call processed response
