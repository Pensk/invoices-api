CREATE TABLE companies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    ceo VARCHAR(255) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    address TEXT
);

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    company_id INT NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id)
);

CREATE TABLE clients (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    ceo VARCHAR(255) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    address TEXT NOT NULL,
    company_id INT NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id)
);

CREATE TABLE client_bank_accounts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    client_id INT NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    branch_name VARCHAR(255) NOT NULL,
    account_number VARCHAR(100) NOT NULL,
    account_name VARCHAR(255) NOT NULL,
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

CREATE TABLE invoices (
    id INT AUTO_INCREMENT PRIMARY KEY,
    issue_date DATE NOT NULL,
    payment_amount BIGINT UNSIGNED NOT NULL,
    fee BIGINT UNSIGNED NOT NULL,
    fee_rate DECIMAL(5, 2) NOT NULL,
    tax BIGINT UNSIGNED NOT NULL,
    tax_rate DECIMAL(5, 2) NOT NULL,
    total_amount BIGINT UNSIGNED NOT NULL,
    due_date DATE NOT NULL,
    status VARCHAR(50) NOT NULL,
    company_id INT NOT NULL,
    client_id INT NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id),
    FOREIGN KEY (client_id) REFERENCES clients(id)
);
