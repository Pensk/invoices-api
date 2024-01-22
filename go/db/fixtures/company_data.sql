INSERT INTO companies (name, ceo, phone, postal_code, address) VALUES
('Company 1', 'CEO 1', '1234567890', '12345', 'Address 1'),
('Company 2', 'CEO 2', '0987654321', '54321', 'Address 2');

SET @company1_id = (SELECT id FROM companies WHERE name = 'Company 1');
SET @company2_id = (SELECT id FROM companies WHERE name = 'Company 2');

INSERT INTO clients (name, ceo, phone, postal_code, address, company_id) VALUES
('Client 1', 'client CEO 1', '1234567890', '12345', 'client Address 1', @company1_id),
('Client 2', 'client CEO 2', '0987654321', '54321', 'client Address 2', @company2_id);

SET @client1_id = (SELECT id FROM clients WHERE name = 'Client 1');
SET @client2_id = (SELECT id FROM clients WHERE name = 'Client 2');

INSERT INTO client_bank_accounts (client_id, bank_name, branch_name, account_number, account_name) VALUES
(@client1_id, 'Bank 1', 'Branch 1', '1234567890', 'client Account 1'),
(@client2_id, 'Bank 2', 'Branch 2', '0987654321', 'client Account 2');