# AI-Driven Code Search Tool for Sourcegraph

## Description

This **AI-Driven Code Search Tool** is designed to enhance Sourcegraph's code intelligence platform by providing a powerful, user-friendly search interface. Built to scale for large, complex repositories and enterprise-scale use cases, it integrates advanced search features such as keyword, regex, and semantic search, allowing users to find code snippets and context with ease. The frontend interface is deployed on GitHub Pages, and the backend is containerized and deployable on Kubernetes, making it a robust and highly adaptable solution.

## Key Features

- **Comprehensive Code Search**: Supports keyword, regex, and semantic code searches.
- **Sourcegraph API Integration**: Uses Sourcegraph’s GraphQL API to fetch and display search results.
- **User-Friendly Interface**: A simple and intuitive React-based UI on GitHub Pages.
- **Enterprise-Ready Backend**: Go-based microservice backend with PostgreSQL and Redis for caching.
- **Automated CI/CD Pipelines**: Uses GitHub Actions for building, testing, and deploying both frontend and backend.
- **Scalable Deployment**: Containerized with Docker and deployable to Kubernetes with Helm charts.

## Table of Contents

- [Getting Started](#getting-started)
- [Setup Instructions](#setup-instructions)
  - [Repository and Files](#repository-and-files)
  - [GitHub Actions](#github-actions)
  - [GitHub Pages](#github-pages)
  - [Kubernetes Deployment](#kubernetes-deployment)
- [Usage Guide](#usage-guide)
- [Troubleshooting and Maintenance](#troubleshooting-and-maintenance)
- [Contributing](#contributing)
- [License](#license)

---

## Getting Started

To get started with this tool, you need a GitHub account and access to Kubernetes (for backend deployment). **No coding experience is required**—just follow the instructions below carefully.

---

## Setup Instructions

### 1. Repository and Files

1. **Create a New GitHub Repository**:
   - Log in to GitHub, create a new repository, and name it (e.g., `AI-Code-Search-Tool`).
   - Choose **Public** (or **Private** if you prefer privacy) and add a README file.

2. **Upload Project Files**:
   - Download or clone the project files.
   - Upload the `frontend`, `backend`, `infra`, and `docs` folders to your GitHub repository.
   - Commit the changes.

### 2. GitHub Actions

This project uses **GitHub Actions** for CI/CD automation. Two workflows will handle testing, building, and deployment.

1. **Create Frontend Workflow**:
   - In your GitHub repository, go to **Actions**, and create a new workflow file at `.github/workflows/frontend.yml`.
   - Paste the following code:

     ```yaml
     name: Frontend CI/CD

     on:
       push:
         branches: [main]

     jobs:
       build:
         runs-on: ubuntu-latest
         steps:
           - uses: actions/checkout@v2
           - uses: actions/setup-node@v2
             with:
               node-version: '16'
           - run: npm install
           - run: npm run build
           - run: npm run deploy
     ```

   - Commit the file.

2. **Create Backend Workflow**:
   - In the `.github/workflows` directory, create another file named `backend.yml` and paste:

     ```yaml
     name: Backend CI/CD

     on:
       push:
         branches: [main]

     jobs:
       build:
         runs-on: ubuntu-latest
         steps:
           - uses: actions/checkout@v2
           - uses: actions/setup-go@v2
             with:
               go-version: '1.18'
           - run: go test ./...
           - name: Build Docker Image
             run: docker build -t ghcr.io/your-repo/backend:latest .
           - name: Push Docker Image
             run: |
               echo "${{ secrets.GITHUB_TOKEN }}" | docker login ghcr.io -u USERNAME --password-stdin
               docker push ghcr.io/your-repo/backend:latest
     ```

   - Commit the file.

### 3. GitHub Pages

1. **Enable GitHub Pages**:
   - Go to **Settings > Pages** in your repository.
   - Under **Source**, select `main` branch and `/docs` folder (or `/frontend/build` if that’s where your build files are).
   - Click **Save**. GitHub will provide a link to your site once deployed.

2. **Access Your Pages URL**:
   - Copy your GitHub Pages URL (e.g., `https://your-username.github.io/AI-Code-Search-Tool`), which will be the frontend link for this tool.

### 4. Kubernetes Deployment

1. **Create a Kubernetes Cluster**:
   - Use a cloud provider (e.g., Google Cloud, AWS, or Azure) to create a Kubernetes cluster.

2. **Deploy Backend**:
   - Open a terminal with `kubectl` (Kubernetes CLI) configured.
   - Run the following commands:

     ```bash
     kubectl apply -f infra/deployment.yaml
     kubectl apply -f infra/service.yaml
     ```

3. **Verify Deployment**:
   - Run `kubectl get pods` and `kubectl get services` to check that your backend is live.

---

## Usage Guide

1. **Visit the Frontend URL**: Go to the GitHub Pages URL provided after deployment.
2. **Run a Search**:
   - Enter a code search term (e.g., a keyword, regex pattern, or semantic query).
   - Choose the type of search (Keyword, Regex, Semantic).
   - Click **Search** to view results.
3. **View Results**: Results are displayed with syntax-highlighted code snippets and links to Sourcegraph for context.

---

## Troubleshooting and Maintenance

If you encounter issues, here are some common checks:

1. **GitHub Actions Failing**:
   - Go to the **Actions** tab to see error logs.
   - Ensure that you have valid API keys or permissions set in your `.env` files.

2. **Backend Issues**:
   - Check Kubernetes logs: `kubectl logs [pod-name]` to identify errors.
   - Make sure `SOURCEGRAPH_API_URL` is correct in your `.env` configuration.

3. **Frontend Display Issues**:
   - Refresh the GitHub Pages URL.
   - Check the **Developer Console** in your browser (press `F12` in most browsers) for errors.

---

## Contributing

Contributions are welcome! Please:
1. Fork this repository.
2. Create a new branch (`feature-branch`).
3. Commit your changes.
4. Open a pull request.

[![Frontend CI/CD](https://github.com/Chrisadams777/AI-Code-Search-Tool/actions/workflows/frontend.yml/badge.svg)](https://github.com/Chrisadams777/AI-Code-Search-Tool/actions/workflows/frontend.yml)

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Credits

Special thanks to the open-source community for providing tools like Sourcegraph, GitHub Actions, Kubernetes, Docker, and Helm that make projects like this possible.
