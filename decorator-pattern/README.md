# Golang Notification System

A flexible notification system implementation using the Decorator pattern in Go. This system demonstrates how to build extensible notification functionality by layering different behaviors such as logging, retry mechanisms, and multi-channel support.

## 📖 Overview

This project showcases the practical implementation of the Decorator design pattern for a notification system. The system allows you to:
- Send notifications through different channels (email, SMS, Slack)
- Add logging capabilities to track notification delivery
- Implement retry mechanisms for failed notifications
- Support multi-channel notification delivery
- Easily extend functionality by adding new decorators

## 🏗️ Project Structure

```
decorator-pattern/
├── main.go                                 # Application entry point
├── notification/
│       ├── notifier.go                     # Core interface definition
│       ├── baseNotifier.go                 # Basic notifier implementation
│       └── decorators/
│           ├── decoratorInterface.go       # Base decorator structure
│           ├── logging.go                  # Logging decorator
│           ├── retry.go                    # Retry mechanism decorator
│           └── multichannel.go             # Multi-channel support decorator
├── go.mod
└── README.md
```

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher

### Installation

1. Clone the repository:
```bash
git clone https://github.com/Ankitcode99/design-patterns-in-go.git
cd decorator-pattern
```


2. Run the application:
```bash
go run main.go
```

## 💡 Usage

### Basic Usage

```go
package main

import (
    "decorator-pattern/internal/notification"
    "decorator-pattern/internal/notification/decorators"
)

func main() {
    // Create a basic notifier
    emailNotifier := notification.NewBasicNotifier("email")

    // Send a simple notification
    emailNotifier.Send("Hello, World!")
}
```

### Adding Decorators

```go
// Create a notifier with logging and retry capabilities
notifier := decorators.NewRetryDecorator(
    decorators.NewLoggingDecorator(emailNotifier),
    3,
)

// Send notification with enhanced capabilities
notifier.Send("Important message")
```

### Multi-channel Notification

```go
// Create different channel notifiers
emailNotifier := notification.NewBasicNotifier("email")
smsNotifier := notification.NewBasicNotifier("SMS")
slackNotifier := notification.NewBasicNotifier("Slack")

// Configure multi-channel notification
notifier := decorators.NewMultiChannelDecorator(
    emailNotifier,
    []notification.Notifier{smsNotifier, slackNotifier},
)

// Send to all channels
notifier.Send("Broadcast message")
```

## 🔧 Available Decorators

1. **LoggingDecorator**
   - Adds timestamp-based logging
   - Tracks success/failure of notifications
   - Usage: `decorators.NewLoggingDecorator(baseNotifier)`

2. **RetryDecorator**
   - Implements retry mechanism with exponential backoff
   - Configurable maximum retry attempts
   - Usage: `decorators.NewRetryDecorator(baseNotifier, maxRetries)`

3. **MultiChannelDecorator**
   - Enables sending to multiple notification channels
   - Supports different channel types
   - Usage: `decorators.NewMultiChannelDecorator(baseNotifier, additionalChannels)`

## 🛠️ Adding New Decorators

To create a new decorator:

1. Create a new file in `internal/notification/decorators/`
2. Define your decorator structure:
```go
type NewDecorator struct {
    BaseDecorator
    // Additional fields if needed
}
```

3. Implement the Notifier interface:
```go
func (d *NewDecorator) Send(message string) error {
    // Add your functionality here
    return d.wrapped.Send(message)
}
```

4. Add a constructor:
```go
func NewNewDecorator(n notification.Notifier) *NewDecorator {
    return &NewDecorator{BaseDecorator: BaseDecorator{wrapped: n}}
}
```

## 📝 Design Decisions

- **Interface-based Design**: Uses Go interfaces for maximum flexibility and extensibility
- **Internal Package**: Core components are in the internal package to prevent external modification
- **Decorator Pattern**: Enables dynamic addition of behaviors without affecting existing code
- **Single Responsibility**: Each decorator handles one specific aspect of notification handling

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/new-decorator`
3. Commit your changes: `git commit -am 'Add new decorator'`
4. Push to the branch: `git push origin feature/new-decorator`
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.