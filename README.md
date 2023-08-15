# Inscriptions Indexer Implementation

This repository delivers an implementation of the indexer following the detailed RFC on Inscriptions Indexers.

## Overview

The Inscriptions Indexer seeks to refine the description protocol by setting stringent standards, aiding in the development process with clearer guidelines supplemented by examples. The innovation revolves around shifting from transaction calldata to "internal" transaction calldata, commonly referred to as trace in the EVM.

## Key Features:

Ethscriptions: Overcomes limitations of the prior ethscription mechanism that relied purely on transaction calldata, thereby preventing interactions with smart contracts.
Internal Transaction Calldata: Uses EVM's trace, making it compatible with both old and new ethscription protocols.
Data URL Encoding: All ethscriptions must follow a valid Data URL encoding to maintain consistency and avoid ambiguities.
Duplication Detection: Provides a structured process for data unification, ensuring that duplicate inscriptions, even if represented differently, are identified effectively.
Usage Guidelines:
DATA URL Encoding: Ensure the Data URL is encoded correctly with reserved characters.

Example:

mathematica
Copy code
is_data_url('data:,{"p":"erc-20","op":"mint","tick":"xxx","id":"10149","amt":"1000"}') == false
is_data_url('data:,%7B%22p%22%3A%22erc-20%22%2C%22op%22%3A%22mint%22%2C%22tick%22%3A%22pepe%22%2C%22id%22%3A%2210149%22%2C%22amt%22%3A%221000%22%7D') == true
Duplication Detection: Follow the series of steps to clean, possibly decode, prepare, and eventually hash the data to ensure its uniqueness.

## Future Developments:

We are exploring the introduction of new mime types specific to tokens and collections. Look out for regular updates and documentation enhancements!

## Contributing

Your contributions are invaluable. Refer to CONTRIBUTING.md for guidelines on making this implementation even better.
