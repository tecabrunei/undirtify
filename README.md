# Undirtify

Undirtify is an open source geotechnical borehole logging tool that allows users to log their boreholes using their laptops or smart devices while on site. This eliminates the need for traditional paper and pen techniques, making the logging process faster, more accurate, and more environmentally friendly. With Undirtify, users can easily record and share data from their boreholes, making it a valuable tool for geotechnical engineers, geologists, and other professionals who work with subsurface data. The open source nature of Undirtify also means that it can be customized and improved by the community, making it a truly collaborative effort. This repository contains the source code for the Undirtify website, built using Golang, PostgreSQL, and React with Tailwind CSS.

## Getting started

### Prerequisites

To run this project, you'll need to have the following installed:

- Docker
- Node.js
- Golang
- PostgreSQL

### Installation

1. Clone this repository:
```sh
git clone https://github.com/tecabrunei/undirtify.git
cd undirtify
```

2. Create a PostgreSQL database:
```sh
docker-compose up -d db
```

3. Create the necessary tables in the database:
```sh
docker-compose run web go run migrate.go
```

4. Install the client dependencies and build the React app:
```sh
cd client
npm install
npm run build
cd ..
```

5. Build and start the Golang server:
```sh
docker-compose up --build
```

6. Access the Undirtify website at http://localhost:8080.

## Usage

The Undirtify website has the following public pages:

- Home
- About Us
- Services
- Careers
- Contact Us

You can use the navigation menu at the top of the page to access these pages.

The Contact Us page has a form that you can use to send a message to the Undirtify team. When you submit the form, your message will be stored in the database.

## Contributing

If you'd like to contribute to this project, please follow these steps:

1. Fork this repository.
2. Create a new branch: `git checkout -b my-new-feature`.
3. Make your changes and commit them: `git commit -am 'Add some feature'`.
4. Push to the branch: `git push origin my-new-feature`.
5. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the LICENSE file for more information.
