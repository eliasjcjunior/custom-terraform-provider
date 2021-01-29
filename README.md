# Terraform Provider JSON
Custom plugin created on article: [Creating the first terraform custom provider](https://aws.amazon.com/pt/sdk-for-node-js/)
   
### Requirements
   * [Golang 1.13+ installed and configured.](https://golang.org/doc/install)
   * [Terraform 0.14+ CLI](https://learn.hashicorp.com/tutorials/terraform/install-cli) installed locally
   * [NodeJS 10+](https://nodejs.org/en/download/)

### Running server
* Install the **json-server** to running a easily and fast server.
    ```shell
    $ npm i -g json-server
    ```
 * On project root create a **db.json** file to initialize the server, with the following content:
    ```
    {
      "users": []
    }
    ```
 * Run the server 
    ```
    $ json-server --watch db.json
    ```

### Build provider
* Run the following command to build the provider on folder **terraform-provider-json**
    ```
    $ go build -o terraform-provider-json
    ```
* Run Makefile to install provider
    ```
    $ make install
    ```
### Running example
* On root folder run terraform with example configuration 
    ```
    $ terraform init && terraform plan
    ```