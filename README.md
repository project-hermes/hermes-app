# hermes-app
[![codebeat badge](https://codebeat.co/badges/37f7b17f-07b8-4346-95f6-badd5be02056)](https://codebeat.co/projects/github-com-project-hermes-hermes-app-master)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/project-hermes/hermes-app.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/project-hermes/hermes-app/alerts/)

# Environments
## Staging
All code that lands on master is auto deployed to the [staging environemtn](https://project-hermes-staging.firebaseapp.com/).

## Prod
All code on the prod branch is auto deployed to [production](https://app.project-hermes.com/) via GCP Cloud Build.

# How to Contribute
## Project
These are based on work in different areas. There is design, back end, and front end (more will be added). The boards are set ip kanban style with automated progression based on the tickets state.

## Issues
Issues are for bugs, features, or just tasks. All work needs to be tracked in an issue.

### Labels
Issues have flags to sort what they are related to, such as design, enhancment etc.

### Assignment
The person the issue is assigned to is responsible for the work of the issue

## Pull Requests
All PRs are to target master. No PR can target prod except from master.

### Reviewers
Each PR must have 1 reviewer that is not the person writing the code.

### Assignees
These people are responsible for merging the card and any ops work needed for the card to land. This person can not be the person the wrote the code.

### Labels
Please use the appropriate labs for the PR. They are used to explain what the PR will relate to.

### Projects
This sets the PR to the proper Kanban board.

### Milestone
These must be set to keep track of current milestone progress.
