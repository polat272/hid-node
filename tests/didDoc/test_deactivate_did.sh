#!bin/sh

first_test_case () {
    echo "----------------Test case 1: A non-controller DID (User-1) trying to deactivate User-2's DID----------------"
    echo "----------------It should throw an Error----------------"
    VERSION_ID=$(hid-noded query ssi did did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52 --output json | jq -r '.didDoc.metadata.versionId')
    hid-noded tx ssi deactivate-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52",
"controller": ["did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52"],
"alsoKnownAs": ["did:hs:1f49341a-de30993e6c52"],
"verificationMethod": [
{
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"type": "Ed25519VerificationKey2020",
"controller": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52",
"publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
}
],
"authentication": [
"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
]
}' "${VERSION_ID}" did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf --ver-key oVtY1xceDZQjkfwlbCEC2vgeADcxpgd27vtYasBhcM/JLR6PnPoD9jvjSJrMsMJwS7faPy5OlFCdj/kgLVZMEg== --from alice --chain-id hidnode --yes
    echo "----------------x----------------"
}

second_test_case () {
    echo "\n\n----------------Test case 2: Deactivating User-2's DID using its verification key----------------"
    VERSION_ID=$(hid-noded query ssi did did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52 --output json | jq -r '.didDoc.metadata.versionId')
    hid-noded tx ssi deactivate-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52",
"controller": ["did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52"],
"alsoKnownAs": ["did:hs:1f49341a-de30993e6c52"],
"verificationMethod": [
{
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"type": "Ed25519VerificationKey2020",
"controller": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52",
"publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
}
],
"authentication": [
"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
]
}' "${VERSION_ID}" did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --ver-key bZBUkLGChnJujYHUZ4L8PECoN2Odv6adWGXc1qVWCRVqtEx0o/FmtFZnd5pT3laR518P58TRUGY5q5KSrToSmQ== --from alice --chain-id hidnode --yes
    echo "----------------x----------------"
}

third_test_case () {
    echo "\n\n----------------Test case 3: Attempting to update the deactivated DID----------------"
    echo "----------------It should throw an Error----------------"
    VERSION_ID=$(hid-noded query ssi did did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52 --output json | jq -r '.didDoc.metadata.versionId')
    hid-noded tx ssi update-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52",
"controller": ["did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52","did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51"],
"alsoKnownAs": ["did:hs:1f49341a-de30993e6c52"],
"verificationMethod": [
{
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"type": "Ed25519VerificationKey2020",
"controller": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52",
"publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
}
],
"authentication": [
"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
]
}' "${VERSION_ID}" did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --ver-key bZBUkLGChnJujYHUZ4L8PECoN2Odv6adWGXc1qVWCRVqtEx0o/FmtFZnd5pT3laR518P58TRUGY5q5KSrToSmQ== --from alice --chain-id hidnode --yes
    echo "----------------x----------------"
}

run_tests() {
    first_test_case
    second_test_case
    third_test_case
}

run_tests