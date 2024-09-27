Feature: Category
  Scenario: User creates a category.
    When I make a request to create a category:
      | title  | Topic Title |
      | points | Key Point 1 |
      | points | Key Point 2 |
    Then I should receive category created success response:
      | title  | Topic Title |
      | points | Key Point 1 |
      | points | Key Point 2 |

  Scenario: User lists categories.
    When I request "/api/category"
    Then I should receive list of all conversation topics
    And  I should see default conversation topics

  Scenario: User updates the category.
    When I request to update previously created category with:
      | title  | Topic Title (new) |
      | points | Key Point 1 (new) |
      | points | Key Point 2       |
    Then category should match:
      | title  | Topic Title (new) |
      | points | Key Point 1 (new) |
      | points | Key Point 2       |

  Scenario: User updates category using only title
    When I request to update previously created category with:
      | title  | Topic Title       |
    Then category should match:
      | title  | Topic Title       |
      | points | Key Point 1 (new) |
      | points | Key Point 2       |

  Scenario: User updates category with invalid title
    When I request to update previously created category using invalid data:
      | title  | ti                |
    Then category should match:
      | title  | Topic Title       |
      | points | Key Point 1 (new) |
      | points | Key Point 2       |

  Scenario: User updates category with invalid title
    When I request to update previously created category using invalid data:
      | title  | ti                |
    Then category should match:
      | title  | Topic Title       |
      | points | Key Point 1 (new) |
      | points | Key Point 2       |

  Scenario: User updates category with invalid point
    When I request to update previously created category using invalid data:
      | title  | Other Title |
      | points | pt          |
    Then category should match:
      | title  | Topic Title       |
      | points | Key Point 1 (new) |
      | points | Key Point 2       |

  Scenario: User deletea the category
    When I request to delete previously created category
    Then category should be unavailable

  Scenario: User can not create a category with too short title
    When I make a request to create a category:
      | title  | ti |
      | points | Key Point 1 |
    Then API returns: category create error: unprocessable entity

  Scenario: User can not create a category with too short point text
    When I make a request to create a category:
      | title  | Topic Title |
      | points | kp |
    Then API returns: category create error: unprocessable entity

  Scenario: User creates a category without points -- Fails
    When I make a request to create a category:
      | title  | Topic Title |
    Then API returns: category create error: unprocessable entity
