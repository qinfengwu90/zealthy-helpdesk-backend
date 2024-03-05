# Helpdesk Ticketing System Backend Server

This repository contains the backend server for a Helpdesk Ticketing System. The system allows 
- **Users**: can submit new tickets, review existing tickets, and receive email updates when the status of their ticket changes. 
- **Admins**: can log in to process tickets by adding admin responses, changing ticket statuses, deleting tickets, creating new admin accounts, and changing passwords.

## Features

- **User Interface**: Provides endpoints for users to interact with the Helpdesk Ticketing System.
- **Ticket Management**: Users can submit new tickets, review existing tickets, and receive email updates on ticket status changes.
- **Admin Panel**: Administrators can log in to process tickets, add admin responses, change ticket statuses, and delete tickets.
- **Authentication**: Secures endpoints with authentication mechanisms for administrators.

## Technologies Used

- **Go**: Backend server environment.
- **PostgreSQL**: SQL database for storing tickets, user info and admin info.
- **JWT (JSON Web Tokens)**: Used for authentication and authorization.

## Deployment
The PostgreSQL database is located on Google Cloud SQL. The Go backend server is hosted using Google Cloud Run.

## Contributing

Contributions are welcome! If you'd like to add features, fix bugs, or improve the documentation, please fork the repository and submit a pull request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

Special thanks to all contributors who have helped to improve this project.

---

Feel free to modify and customize this README according to your specific project requirements. Happy coding! 😊
