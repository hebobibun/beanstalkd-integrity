# Beanstalkd Data Integrity Testing

This repository contains a simple setup to test data integrity using [Beanstalkd](https://beanstalkd.github.io/), a simple, fast work queue service.

## Overview

Beanstalkd is a lightweight messaging system that provides reliable message delivery. This repository demonstrates how to use Beanstalkd to test data integrity by publishing and subscribing to messages.

## Setup

1. **Install Beanstalkd**: If you haven't already, install Beanstalkd on your system. You can find installation instructions in the [Beanstalkd documentation](https://beanstalkd.github.io/download.html).

2. **Clone the Repository**: Clone this repository to your local machine using Git:

    ```bash
    git clone https://github.com/hebobibun/beanstalkd-integrity.git
    ```

## Testing Data Integrity

1. **Run Publisher**: Start the publisher service to publish messages to Beanstalkd.

    ```bash
    cd pub
    go run main.go
    ```

2. **Run Subscriber 1**: Start the subscriber service to receive and process messages from Beanstalkd.

    ```bash
    cd sub1
    go run main.go
    ```

3. **Run Subscriber 2**: Start the subscriber service to receive and process messages from Beanstalkd.

    ```bash
    cd sub2
    go run main.go
    ```

4. **Verify Integrity**: Monitor the subscriber output to ensure that messages are received intact and in the correct order.

## Contributing

Contributions are welcome! If you find issues or have ideas for improvements related to testing data integrity with Beanstalkd, feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
