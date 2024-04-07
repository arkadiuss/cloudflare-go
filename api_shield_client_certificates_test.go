package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListHostnameAssociations(t *testing.T) {
	setup()
	defer teardown()

	params := ListHostnameAssociationParams{
		CertificateID: "af7149d5-1ca0-4768-8bb1-d50a51c7d845",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"mtls_certificate_id": "af7149d5-1ca0-4768-8bb1-d50a51c7d845",
				"hostnames": [
					"test.example.com"
				]
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/certificate_authorities/hostname_associations", handler)

	want := HostnameAssociation{
		CertificateID: "af7149d5-1ca0-4768-8bb1-d50a51c7d845",
		Hostnames:     []string{"test.example.com"},
	}

	actual, err := client.ListHostnameAssociations(context.Background(), ZoneIdentifier(testZoneID), params)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestListHostnameAssociationsWithoutID(t *testing.T) {
	setup()
	defer teardown()

	params := ListHostnameAssociationParams{}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"hostnames": [
					"test.example.com"
				]
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/certificate_authorities/hostname_associations", handler)

	want := HostnameAssociation{
		Hostnames: []string{"test.example.com"},
	}

	actual, err := client.ListHostnameAssociations(context.Background(), ZoneIdentifier(testZoneID), params)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestReplaceHostnameAssociations(t *testing.T) {
	setup()
	defer teardown()

	params := ReplaceHostnameAssociationParams{
		CertificateID: "52ad7c3e-f37e-4835-8b24-07d8afcaeb3",
		Hostnames:     []string{"test.example.com"},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"mtls_certificate_id": "52ad7c3e-f37e-4835-8b24-07d8afcaeb3",
				"hostnames": [
					"test.example.com"
				]
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/certificate_authorities/hostname_associations", handler)

	want := HostnameAssociation{
		CertificateID: "52ad7c3e-f37e-4835-8b24-07d8afcaeb3",
		Hostnames:     []string{"test.example.com"},
	}

	actual, err := client.ReplaceHostnameAssociations(context.Background(), ZoneIdentifier(testZoneID), params)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestListClientCertificates(t *testing.T) {
	setup()
	defer teardown()

	params := ListClientCertificatesParams{}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [{
				"certificate": "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoC...dhDDE\n-----END CERTIFICATE-----",
				"certificate_authority": {
					"id": "568b6b74-7b0c-4755-8840-4e3b8c24adeb",
					"name": "Cloudflare Managed CA for account"
				},
				"common_name": "Cloudflare",
				"country": "US",
				"csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICY....\n-----END CERTIFICATE REQUEST-----\n",
				"expires_on": "2033-02-20T23:18:00Z",
				"fingerprint_sha256": "256c24690243359fb8cf139a125bd05ebf1d968b71e4caf330718e9f5c8a89ea",
				"id": "023e105f4ecef8ad9ca31a8372d0c353",
				"issued_on": "2023-02-23T23:18:00Z",
				"location": "Somewhere",
				"organization": "Organization",
				"organizational_unit": "Organizational Unit",
				"serial_number": "3bb94ff144ac567b9f75ad664b6c55f8d5e48182",
				"signature": "SHA256WithRSA",
				"ski": "8e375af1389a069a0f921f8cc8e1eb12d784b949",
				"state": "CA",
				"status": "active",
				"validity_days": 3650
			}]
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/client_certificates", handler)

	want := []ClientCertificate{
		{
			Certificate: "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoC...dhDDE\n-----END CERTIFICATE-----",
			CertificateAuthority: ClientCertificateAuthority{
				ID:   "568b6b74-7b0c-4755-8840-4e3b8c24adeb",
				Name: "Cloudflare Managed CA for account",
			},
			CommonName:         "Cloudflare",
			Country:            "US",
			CSR:                "-----BEGIN CERTIFICATE REQUEST-----\nMIICY....\n-----END CERTIFICATE REQUEST-----\n",
			ExpiresOn:          "2033-02-20T23:18:00Z",
			FingerprintSHA256:  "256c24690243359fb8cf139a125bd05ebf1d968b71e4caf330718e9f5c8a89ea",
			ID:                 "023e105f4ecef8ad9ca31a8372d0c353",
			IssuedOn:           "2023-02-23T23:18:00Z",
			Location:           "Somewhere",
			Organization:       "Organization",
			OrganizationalUnit: "Organizational Unit",
			SerialNumber:       "3bb94ff144ac567b9f75ad664b6c55f8d5e48182",
			Signature:          "SHA256WithRSA",
			SKI:                "8e375af1389a069a0f921f8cc8e1eb12d784b949",
			State:              "CA",
			Status:             "active",
			ValidityDays:       3650,
		},
	}

	actual, _, err := client.ListClientCertificates(context.Background(), ZoneIdentifier(testZoneID), params)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateClientCertificate(t *testing.T) {
	setup()
	defer teardown()

	params := CreateClientCertificateParams{
		ValidityDays: 365,
		CSR:          "-----BEGIN CERTIFICATE REQUEST-----\nMIICY....\n-----END CERTIFICATE REQUEST-----",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"certificate": "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoC...dhDDE\n-----END CERTIFICATE-----",
				"certificate_authority": {
					"id": "568b6b74-7b0c-4755-8840-4e3b8c24adeb",
					"name": "Cloudflare Managed CA for account"
				},
				"common_name": "Cloudflare",
				"country": "US",
				"csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICY....\n-----END CERTIFICATE REQUEST-----\n",
				"expires_on": "2033-02-20T23:18:00Z",
				"fingerprint_sha256": "256c24690243359fb8cf139a125bd05ebf1d968b71e4caf330718e9f5c8a89ea",
				"id": "023e105f4ecef8ad9ca31a8372d0c353",
				"issued_on": "2023-02-23T23:18:00Z",
				"location": "Somewhere",
				"organization": "Organization",
				"organizational_unit": "Organizational Unit",
				"serial_number": "3bb94ff144ac567b9f75ad664b6c55f8d5e48182",
				"signature": "SHA256WithRSA",
				"ski": "8e375af1389a069a0f921f8cc8e1eb12d784b949",
				"state": "CA",
				"status": "active",
				"validity_days": 365
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/client_certificates", handler)

	want := ClientCertificate{
		Certificate: "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoC...dhDDE\n-----END CERTIFICATE-----",
		CertificateAuthority: ClientCertificateAuthority{
			ID:   "568b6b74-7b0c-4755-8840-4e3b8c24adeb",
			Name: "Cloudflare Managed CA for account",
		},
		CommonName:         "Cloudflare",
		Country:            "US",
		CSR:                "-----BEGIN CERTIFICATE REQUEST-----\nMIICY....\n-----END CERTIFICATE REQUEST-----\n",
		ExpiresOn:          "2033-02-20T23:18:00Z",
		FingerprintSHA256:  "256c24690243359fb8cf139a125bd05ebf1d968b71e4caf330718e9f5c8a89ea",
		ID:                 "023e105f4ecef8ad9ca31a8372d0c353",
		IssuedOn:           "2023-02-23T23:18:00Z",
		Location:           "Somewhere",
		Organization:       "Organization",
		OrganizationalUnit: "Organizational Unit",
		SerialNumber:       "3bb94ff144ac567b9f75ad664b6c55f8d5e48182",
		Signature:          "SHA256WithRSA",
		SKI:                "8e375af1389a069a0f921f8cc8e1eb12d784b949",
		State:              "CA",
		Status:             "active",
		ValidityDays:       365,
	}

	actual, err := client.CreateClientCertificate(context.Background(), ZoneIdentifier(testZoneID), params)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetClientCertificateDetails(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"certificate": "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoC...dhDDE\n-----END CERTIFICATE-----",
				"certificate_authority": {
					"id": "568b6b74-7b0c-4755-8840-4e3b8c24adeb",
					"name": "Cloudflare Managed CA for account"
				},
				"common_name": "Cloudflare",
				"country": "US",
				"csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICY....\n-----END CERTIFICATE REQUEST-----\n",
				"expires_on": "2033-02-20T23:18:00Z",
				"fingerprint_sha256": "256c24690243359fb8cf139a125bd05ebf1d968b71e4caf330718e9f5c8a89ea",
				"id": "d8b1e36d-eff0-4f89-ad8a-1ffc51a08954",
				"issued_on": "2023-02-23T23:18:00Z",
				"location": "Somewhere",
				"organization": "Organization",
				"organizational_unit": "Organizational Unit",
				"serial_number": "3bb94ff144ac567b9f75ad664b6c55f8d5e48182",
				"signature": "SHA256WithRSA",
				"ski": "8e375af1389a069a0f921f8cc8e1eb12d784b949",
				"state": "CA",
				"status": "active",
				"validity_days": 365
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/client_certificates/d8b1e36d-eff0-4f89-ad8a-1ffc51a08954", handler)

	want := ClientCertificate{
		Certificate: "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoC...dhDDE\n-----END CERTIFICATE-----",
		CertificateAuthority: ClientCertificateAuthority{
			ID:   "568b6b74-7b0c-4755-8840-4e3b8c24adeb",
			Name: "Cloudflare Managed CA for account",
		},
		CommonName:         "Cloudflare",
		Country:            "US",
		CSR:                "-----BEGIN CERTIFICATE REQUEST-----\nMIICY....\n-----END CERTIFICATE REQUEST-----\n",
		ExpiresOn:          "2033-02-20T23:18:00Z",
		FingerprintSHA256:  "256c24690243359fb8cf139a125bd05ebf1d968b71e4caf330718e9f5c8a89ea",
		ID:                 "d8b1e36d-eff0-4f89-ad8a-1ffc51a08954",
		IssuedOn:           "2023-02-23T23:18:00Z",
		Location:           "Somewhere",
		Organization:       "Organization",
		OrganizationalUnit: "Organizational Unit",
		SerialNumber:       "3bb94ff144ac567b9f75ad664b6c55f8d5e48182",
		Signature:          "SHA256WithRSA",
		SKI:                "8e375af1389a069a0f921f8cc8e1eb12d784b949",
		State:              "CA",
		Status:             "active",
		ValidityDays:       365,
	}

	actual, err := client.GetClientCertificateDetails(context.Background(), ZoneIdentifier(testZoneID), "d8b1e36d-eff0-4f89-ad8a-1ffc51a08954")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestRevokeClientCertificate(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"certificate": "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoC...dhDDE\n-----END CERTIFICATE-----",
				"certificate_authority": {
					"id": "568b6b74-7b0c-4755-8840-4e3b8c24adeb",
					"name": "Cloudflare Managed CA for account"
				},
				"common_name": "Cloudflare",
				"country": "US",
				"csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICY....\n-----END CERTIFICATE REQUEST-----\n",
				"expires_on": "2033-02-20T23:18:00Z",
				"fingerprint_sha256": "256c24690243359fb8cf139a125bd05ebf1d968b71e4caf330718e9f5c8a89ea",
				"id": "1391a76b-7cdf-4fef-a7b0-c326bf39cfc8",
				"issued_on": "2023-02-23T23:18:00Z",
				"location": "Somewhere",
				"organization": "Organization",
				"organizational_unit": "Organizational Unit",
				"serial_number": "3bb94ff144ac567b9f75ad664b6c55f8d5e48182",
				"signature": "SHA256WithRSA",
				"ski": "8e375af1389a069a0f921f8cc8e1eb12d784b949",
				"state": "CA",
				"status": "pending_revocation",
				"validity_days": 365
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/client_certificates/1391a76b-7cdf-4fef-a7b0-c326bf39cfc8", handler)

	want := ClientCertificate{
		Certificate: "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoC...dhDDE\n-----END CERTIFICATE-----",
		CertificateAuthority: ClientCertificateAuthority{
			ID:   "568b6b74-7b0c-4755-8840-4e3b8c24adeb",
			Name: "Cloudflare Managed CA for account",
		},
		CommonName:         "Cloudflare",
		Country:            "US",
		CSR:                "-----BEGIN CERTIFICATE REQUEST-----\nMIICY....\n-----END CERTIFICATE REQUEST-----\n",
		ExpiresOn:          "2033-02-20T23:18:00Z",
		FingerprintSHA256:  "256c24690243359fb8cf139a125bd05ebf1d968b71e4caf330718e9f5c8a89ea",
		ID:                 "1391a76b-7cdf-4fef-a7b0-c326bf39cfc8",
		IssuedOn:           "2023-02-23T23:18:00Z",
		Location:           "Somewhere",
		Organization:       "Organization",
		OrganizationalUnit: "Organizational Unit",
		SerialNumber:       "3bb94ff144ac567b9f75ad664b6c55f8d5e48182",
		Signature:          "SHA256WithRSA",
		SKI:                "8e375af1389a069a0f921f8cc8e1eb12d784b949",
		State:              "CA",
		Status:             "pending_revocation",
		ValidityDays:       365,
	}

	actual, err := client.RevokeClientCertificate(context.Background(), ZoneIdentifier(testZoneID), "1391a76b-7cdf-4fef-a7b0-c326bf39cfc8")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestReactivateClientCertificate(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"certificate": "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoC...dhDDE\n-----END CERTIFICATE-----",
				"certificate_authority": {
					"id": "568b6b74-7b0c-4755-8840-4e3b8c24adeb",
					"name": "Cloudflare Managed CA for account"
				},
				"common_name": "Cloudflare",
				"country": "US",
				"csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICY....\n-----END CERTIFICATE REQUEST-----\n",
				"expires_on": "2033-02-20T23:18:00Z",
				"fingerprint_sha256": "256c24690243359fb8cf139a125bd05ebf1d968b71e4caf330718e9f5c8a89ea",
				"id": "2ebcb1e6-4c13-4f0c-9be6-1fd4767b543e",
				"issued_on": "2023-02-23T23:18:00Z",
				"location": "Somewhere",
				"organization": "Organization",
				"organizational_unit": "Organizational Unit",
				"serial_number": "3bb94ff144ac567b9f75ad664b6c55f8d5e48182",
				"signature": "SHA256WithRSA",
				"ski": "8e375af1389a069a0f921f8cc8e1eb12d784b949",
				"state": "CA",
				"status": "active",
				"validity_days": 365
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/client_certificates/2ebcb1e6-4c13-4f0c-9be6-1fd4767b543e", handler)

	want := ClientCertificate{
		Certificate: "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoC...dhDDE\n-----END CERTIFICATE-----",
		CertificateAuthority: ClientCertificateAuthority{
			ID:   "568b6b74-7b0c-4755-8840-4e3b8c24adeb",
			Name: "Cloudflare Managed CA for account",
		},
		CommonName:         "Cloudflare",
		Country:            "US",
		CSR:                "-----BEGIN CERTIFICATE REQUEST-----\nMIICY....\n-----END CERTIFICATE REQUEST-----\n",
		ExpiresOn:          "2033-02-20T23:18:00Z",
		FingerprintSHA256:  "256c24690243359fb8cf139a125bd05ebf1d968b71e4caf330718e9f5c8a89ea",
		ID:                 "2ebcb1e6-4c13-4f0c-9be6-1fd4767b543e",
		IssuedOn:           "2023-02-23T23:18:00Z",
		Location:           "Somewhere",
		Organization:       "Organization",
		OrganizationalUnit: "Organizational Unit",
		SerialNumber:       "3bb94ff144ac567b9f75ad664b6c55f8d5e48182",
		Signature:          "SHA256WithRSA",
		SKI:                "8e375af1389a069a0f921f8cc8e1eb12d784b949",
		State:              "CA",
		Status:             "active",
		ValidityDays:       365,
	}

	actual, err := client.ReactivateClientCertificate(context.Background(), ZoneIdentifier(testZoneID), "2ebcb1e6-4c13-4f0c-9be6-1fd4767b543e")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
