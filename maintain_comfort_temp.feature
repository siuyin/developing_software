Feature: Maintain a comfortable temperature
	As a home resident
	I want to have a comfortable temperature
	that is not too hot nor too cold.

	Scenario: Cooling when too hot
		Given a Thermometer
		When Thermometer reads too hot
		Then I should feel cooling effects

	Scenario: Heating when too cold
		Given a Thermometer
		When Thermometer reads too cold
		Then I should feel heating effects
