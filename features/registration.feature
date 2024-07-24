Feature: OAuth 2.0 Registration
  As a user
  I want to register using OAuth 2.0
  So that I can securely use the application with my existing credentials

  Background:
    Given the OAuth 2.0 service is available

  Scenario: Successful registration with OAuth 2.0
    Given I have chosen to register using an OAuth 2.0 provider
    When I select my preferred OAuth provider
    And I authorize the application to access my information
    Then I should be redirected back to the application
    And I should be registered and logged in successfully

  Scenario: Failed registration due to OAuth 2.0 service unavailability
    Given the OAuth 2.0 service is unavailable
    When I attempt to register using an OAuth 2.0 provider
    Then I should see an error message indicating the service is unavailable

  Scenario: User cancels the OAuth 2.0 authorization
    Given I have chosen to register using an OAuth 2.0 provider
    When I cancel the authorization request
    Then I should be redirected back to the registration page
    And I should see a message indicating that the registration was not completed