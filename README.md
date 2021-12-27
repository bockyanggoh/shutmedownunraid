# Shut Me Down Unraid
A full stack application with UI + API Server to communicate
with Unraid to shutdown unraid server.

The following features will be planned for cycle 1 development.
- UI based scheduling for shutdowns
- On demand shutdown that can work with homekit

# How does it shutdown?
This application is meant to be run as a Docker container, which will
communicate via CLI or API to the host Unraid machine to safely shutdown Unraid Server.