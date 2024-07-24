Feature: OAuth 2.0 Authentication
  As a user
  I want to authenticate using OAuth 2.0
  So that I can securely access the application with my existing credentials

  Background:
    Given the OAuth 2.0 service is available

  Scenario: Successful authentication with OAuth 2.0
    Given I have chosen to authenticate using an OAuth 2.0 provider
    When I select my preferred OAuth provider
    And I authorize the application to access my information
    Then I should be redirected back to the application
    And I should be logged in successfully

  Scenario: Failed authentication due to OAuth 2.0 service unavailability
    Given the OAuth 2.0 service is unavailable
    When I attempt to authenticate using an OAuth 2.0 provider
    Then I should see an error message indicating the service is unavailable

  Scenario: User cancels the OAuth 2.0 authorization
    Given I have chosen to authenticate using an OAuth 2.0 provider
    When I cancel the authorization request
    Then I should be redirected back to the login page
    And I should see a message indicating that the authentication was not completed

  Scenario: Authentication with expired OAuth 2.0 token
    Given I have an expired OAuth 2.0 token
    When I attempt to authenticate
    Then I should be prompted to re-authorize the application
