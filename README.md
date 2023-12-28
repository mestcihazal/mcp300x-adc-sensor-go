# mcp3004-8-go
The module provides analog-to-digital conversion capabilities for MCP3004 and MCP3008 SPI ADCs. Written in Go, board agnostic. Tested on Raspberry Pi and Jetson Orin. 

The values you get from the readings will be proportional to the voltage and you will need to interpret it depending on which sensor you hook up to the MCP300x. 

To use this module in your smart machine configuration, deploy the module to your machine or add it as a local module with the proper executable path. 

Then add the sensor (either from the registry or as a local component) by clicking on the "Create component" button. 

Example configuration for the Attributes section of your sensor looks like this:

{
  "chip_select": "0",
  "spi_bus": "0", 
  "pins": {
    "moisture": 0,
    "temperature": 1,
    "humidity": 2
  }
}

You will need to tell the sensor which chip_select pin you are using as well as the spi_bus. 

You can add as many analog sensors as your MCP300x allows and get readings from them concurrently (this depends on how many channels it has, so for MCP3004 that is four channels, and for MCP3008 that is eight). 
