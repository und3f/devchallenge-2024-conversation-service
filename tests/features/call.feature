Feature: Call
  Scenario: Create call
    When I make a request to create sample call
    Then I should receive call created success response
    And  get call should return accepted response

  Scenario: Get call
    When I make a request to get sample processed call
    And  I should receive call processed response
