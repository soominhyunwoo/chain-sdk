<!--
order: 1
-->

# Concepts

## Authorization and Grant
`x/authz` module defines interfaces and messages grant authorizations to perform actions
on behalf of one account to other accounts. The design is defined in the [ADR 030](../../../architecture/adr-030-authz-module.md).

Grant is an allowance to execute an Msg by grantee address on behalf of the granter.
Authorization is an interface which must be implemented by a concrete authorization logic to validate and execute grants. They are extensible and can be defined for any Msg service method even outside of the module where the Msg method is defined. See the `SendAuthorization` example below for more details.


+++ https://github.com/soominhyunwoo/chain-sdk/blob/c95de9c4177442dee4c69d96917efc955b5d19d9/x/authz/types/authorizations.go#L15-L24


## Built-in Authorizations

soominhyunwoo-SDK `x/authz` module comes with following authorization types

### SendAuthorization

`SendAuthorization` implements `Authorization` interface for the `soominhyunwoo.bank.v1beta1.MsgSend` Msg, that takes a `SpendLimit` and updates it down to zero.

+++ https://github.com/soominhyunwoo/chain-sdk/blob/c95de9c4177442dee4c69d96917efc955b5d19d9/proto/soominhyunwoo/authz/v1beta1/authz.proto#L12-L19

+++ https://github.com/soominhyunwoo/chain-sdk/blob/c95de9c4177442dee4c69d96917efc955b5d19d9/x/authz/types/send_authorization.go#L23-L45

- `spent_limit` keeps track of how many coins left in the authorization.


### GenericAuthorization

`GenericAuthorization` implements the `Authorization` interface, that gives unrestricted permission to execute the provided Msg on behalf of granter's account.

+++ https://github.com/soominhyunwoo/chain-sdk/blob/c95de9c4177442dee4c69d96917efc955b5d19d9/proto/soominhyunwoo/authz/v1beta1/authz.proto#L21-L30

+++ https://github.com/soominhyunwoo/chain-sdk/blob/c95de9c4177442dee4c69d96917efc955b5d19d9/x/authz/types/generic_authorization.go#L20-L28

- `msg` stores Msg type URL.

## Gas

In order to prevent DoS attacks, granting `StakeAuthorizaiton`s with `x/authz` incur gas. `StakeAuthorizaiton` allows you to authorize another account to delegate, undelegate, or redelegate to validators. The authorizer can define a list of validators they will allow and/or deny delegations to. The SDK will iterate over these lists and charge 10 gas for each validator in both of the lists.
