# k8s-deploy-scaffold

A CLI tool to scaffold Kubernetes namespaces and service accounts, with optional AWS IAM Role for Service Accounts (IRSA) integration, based on a YAML configuration file 🚀 

## Features

- Create Kubernetes namespaces and service accounts from a YAML config
- Optionally create AWS IAM roles and annotate service accounts for IRSA
- Supports custom kubeconfig and AWS profile selection

## Prerequisites

- Go 1.24+
- Access to a Kubernetes cluster (kubeconfig)
- AWS credentials (for IRSA/IAM role creation)
- [kubectl](https://kubernetes.io/docs/tasks/tools/) (optional, for manual verification)

## Usage

> **Note:** Deletions are not currently supported. The tool only creates resources as specified in your config and does not save or track state.

1. **Download the latest release:**
   - Go to the [Releases](https://github.com/ramalabeysekera/k8s-deploy-scaffold/releases) page.
   - Download the appropriate binary for your platform.
   - Make it executable if needed:
     ```sh
     chmod +x k8s-deploy-scaffold
     ```

2. **Configure your `config.yml`:**
   See the example below for structure.

3. **Run the tool:**
   ```sh
   ./k8s-deploy-scaffold -kubeconfig /path/to/kubeconfig --verbose
   ```
   - `-kubeconfig` (optional): Path to your kubeconfig file (defaults to `~/.kube/config`)
   - `-verbose`: Enable verbose output

## Development

If you want to build from source:

1. **Clone the repository:**
   ```sh
   git clone https://github.com/ramalabeysekera/k8s-deploy-scaffold.git
   cd k8s-deploy-scaffold
   ```

2. **Build the CLI:**
   ```sh
   go build -o k8s-deploy-scaffold
   ```

## Configuration

Edit `config.yml` to define your workloads. Example:

```yaml
testWorkload:
  - namespace: test-ns
    create_namespace: true
    create_service_account: true
    service_account_name: serviceaccount1
    enable_irsa: true
    oidc_provider: oidc.eks.ap-southeast-1.amazonaws.com/id/XXXXXXXXXXXXXXXXXXXXXXXX
testWorkload2:
  - namespace: test-ns2
    create_namespace: true
    create_service_account: true
    service_account_name: serviceaccount2
    enable_irsa: false
    oidc_provider: oidc.eks.ap-northeast-1.amazonaws.com/id/XXXXXXXXXXXXXXXXXXXXXXXX
```

- `namespace`: Name of the Kubernetes namespace.
- `create_namespace`: Whether to create the namespace.
- `create_service_account`: Whether to create the service account.
- `service_account_name`: Name of the service account.
- `enable_irsa`: If true, creates an AWS IAM role and annotates the service account for IRSA.
- `oidc_provider`: OIDC provider URL for your EKS cluster (required for IRSA).

## AWS Credentials

The tool uses the `aws_profile` environment variable to select the AWS profile (defaults to `default`).

```sh
export aws_profile=my-aws-profile
```

## License

[Apache 2.0](LICENSE)