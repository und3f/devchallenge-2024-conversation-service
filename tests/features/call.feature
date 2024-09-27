Feature: Call

  Scenario: User get processed call details
    When I make a request to get a sample processed call
    And  I should receive call processed response:
      | id             | 1000 |
      | text           | Hello and welcome to out call in Kyiv. I am happy to talk about visa and diplomatic inquries! |
      | name           | Sample Call |
      | location       | Kyiv |
      | emotional_tone | Neutral |
      | categories     | Diplomatic Inquiries |
      | categories     | Visa and Passport Services |

  Scenario: User tries to create a call with invalid audio document
    When I make a request to create a call
      | audio_url | https://example.com/not-audio-file.txt |
    Then I should receive call created success response
    And  I wait till the call is processed
    And  get call should return unprocessable entity

  Scenario: Get non-existing call
    When I make a request to get non-existing call id
    Then I should receive not found error

  @long
  Scenario: User successfully analyzes call
    When I make a request to create a call
      | audio_url | https://github.com/ggerganov/whisper.cpp/raw/refs/heads/master/samples/jfk.wav |
    Then I should receive call created success response
    And  I wait till the call is processed using long poll
    And  get call should return success response

