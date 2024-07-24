Feature: OAuth 2.0 Authorization
  As a user
  I want to authorize applications using OAuth 2.0
  So that I can securely share my data with third-party applications without exposing my credentials

  Background:
    Given the OAuth 2.0 service is available

  Scenario: Successful authorization with OAuth 2.0
    Given I have logged in to the OAuth 2.0 provider
    When I authorize a third-party application
    Then I should receive an authorization code
    And the third-party application should receive an access token

  Scenario: Failed authorization due to OAuth 2.0 service unavailability
    Given the OAuth 2.0 service is unavailable
    When I attempt to authorize a third-party application
    Then I should see an error message indicating the service is unavailable

  Scenario: User denies authorization to the third-party application
    Given I have logged in to the OAuth 2.0 provider
    When I deny the authorization request from a third-party application
    Then the third-party application should not receive an access token
    And I should be redirected back with an error message indicating the denial

  Scenario: Authorization with expired access token
    Given I have an expired access token for a third-party application
    When the third-party application attempts to access my resources
    Then it should receive an unauthorized response
    And I should be prompted to re-authorize the application
