# Tracking Github Events Activities

## Brief Description

This program is a CLI (Command Line Interface) Tool accessible using the operating system terminal environment, it provides information about a github user's events activities you can filter based on the github available events [https://docs.github.com/en/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28](https://docs.github.com/en/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28)

# Getting Started

#### Project URL:

[https://roadmap.sh/projects/github-user-activity](https://roadmap.sh/projects/github-user-activity)

## Prerequisites

- You need to have go installed on your computer to run the project
- You can also build and install the project using this source code code samples would be down below

## Installation

1. Clone the repository
   ```
        git clone <repository_url>
   ```
2. Run the program
   ```
        go run . <github_username>
        go run . <github_username> <eventType>
   ```

## Build

1. Clone the repository
   ```
        git clone <repository_url>
   ```
2. Build
   ```
        go build
        go build -o <custom_program_name>
   ```
3. Install the app
   ```
       go install
   ```
4. Run the built and installed app
   ```
       <directory_name> <github_username>
       <directory_name> <github_username> <event_type>
       <custom_program_name> <github_username>
       <custom_program_name> <github_username> <event_type>
   ```

## Build for different OS

1. WINDOWS
   ```
       GOOS=windows GOARCH=amd64 go build -o <myapp.exe>
   ```
2. LINUX
   ```
       GOOS=linux GOARCH=amd64 go build -o <myapp>
   ```
3. MAC
   ```
       GOOS=darwin GOARCH=arm64 go build -o <myapp>
   ```

# Contributing

Contributions are welcome! If you'd like to contribute, please fork the repository and create a pull request with your changes.

---

# License

Task-Tracker is open-source software licensed under the GNU General Public License Version 3 or any later version. You can find the license [here](https://www.gnu.org/licenses/gpl-3.0.html).

---

# Contact

For questions, feedback, or support, feel free to reach out via email:

- [bidemi64@gmail.com](mailto:bidemi64@gmail.com)
