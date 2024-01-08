# Clean Architecture Module Generator

## About

This repository hosts a code generator written in Go, designed to produce clean architecture template code. The generator automates the creation of key components based on YAML configurations, fostering the development of applications with a clear and maintainable structure.

### Key Features

- **Clean Architecture Templates:** Generate foundational code components following the principles of clean architecture.
- **Domain Models:** Automatically create Go structs for domain models based on YAML specifications.
- **Interfaces:** Generate repository interfaces with predefined methods, promoting separation of concerns.
- **Schema Definitions:** Create schema models compatible with MongoDB to support clean data storage.

### Motivation

The motivation behind this project is to streamline the development workflow by eliminating repetitive code-writing tasks. It aims to enhance code consistency and reduce the chances of human errors in creating foundational components.

## Clean Architecture Template

This code generator provides a Clean Architecture template that follows best practices for structuring and organizing your Go project. The generated code is designed to facilitate maintainability, scalability, and testability.

### Key Characteristics

- **Separation of Concerns:** The codebase is organized into layers, including the core business logic, interfaces, and external dependencies, ensuring a clear separation of concerns.

- **Dependency Rule:** Dependencies flow from outer layers toward inner layers, adhering to the dependency inversion principle. This allows for flexibility in choosing implementations while maintaining a stable core.

- **Testability:** The architecture promotes unit testing by keeping business logic isolated from external dependencies. Interfaces and dependency injection facilitate easy testing and mocking.

- **Flexibility in Infrastructure:** The infrastructure layer is abstracted, allowing for easy replacement of external components such as databases or third-party APIs without affecting the core business logic.

### Project Structure

The generated project structure follows a clean and modular layout, providing dedicated directories for each layer of the architecture.

- `cmd`: Entry point for the application and setup of the dependency injection.
- `domain`: Contains the core business logic, including domain models and use cases.
- `repository`: Houses repository interfaces and implementations.
- `infrastructure`: Manages external dependencies, such as database connections and third-party integrations.

### How to Use

1. **Generate Code:** Utilize the code generator to create a new project or update an existing one based on your YAML configuration.

2. **Review Generated Code:** Examine the generated files in the respective directories to understand how the Clean Architecture principles are implemented.

3. **Customization:** Tailor the generated code to fit your specific business requirements while keeping the Clean Architecture principles in mind.

### Example YAML Configuration

```yaml
# Example YAML configuration for generating Clean Architecture template
module: myapp
file_name: main.go
models:
  - name: User
    fields:
      - name: ID
        type: string
        optional: false
      - name: Name
        type: string
        optional: false
interfaces:
  - name: UserRepository
    methods:
      - name: GetByID
        returns:
          - "*domain.User"
          - "error"
        params:
          - name: "id"
            type: "string"
```

### How it Works

The generator reads YAML configuration files located in the `configs` directory, interprets the specifications and produces well-structured Go code in designated project folders. The generated code follows the principles of clean architecture, ensuring a separation of concerns and maintainability.

### Getting Started

To start using the clean architecture template code generator, follow the instructions in the [Getting Started](#getting-started) section of this README.

### Contribution

Contributions to this project are highly encouraged! Whether it's bug fixes, feature enhancements, or suggestions, feel free to open issues and submit pull requests. Your involvement is crucial in making this tool
