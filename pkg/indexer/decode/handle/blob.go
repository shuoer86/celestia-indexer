// SPDX-FileCopyrightText: 2023 PK Lab AG <contact@pklab.io>
// SPDX-License-Identifier: MIT

package handle

import (
	"github.com/celenium-io/celestia-indexer/internal/storage"
	storageTypes "github.com/celenium-io/celestia-indexer/internal/storage/types"
	"github.com/celenium-io/celestia-indexer/pkg/types"
	"github.com/celestiaorg/celestia-app/pkg/namespace"
	appBlobTypes "github.com/celestiaorg/celestia-app/x/blob/types"
	"github.com/pkg/errors"
)

// MsgPayForBlobs pays for the inclusion of a blob in the block.
func MsgPayForBlobs(level types.Level, m *appBlobTypes.MsgPayForBlobs) (storageTypes.MsgType, []storage.AddressWithType, []storage.Namespace, int64, error) {
	var blobsSize int64
	namespaces := make([]storage.Namespace, len(m.Namespaces))

	for nsI, ns := range m.Namespaces {
		if len(m.BlobSizes) < nsI {
			return storageTypes.MsgUnknown, nil, nil, 0, errors.Errorf(
				"blob sizes length=%d is less then namespaces index=%d", len(m.BlobSizes), nsI)
		}

		appNS := namespace.Namespace{Version: ns[0], ID: ns[1:]}
		size := int64(m.BlobSizes[nsI])
		blobsSize += size
		namespaces[nsI] = storage.Namespace{
			FirstHeight: level,
			Version:     appNS.Version,
			NamespaceID: appNS.ID,
			Size:        size,
			PfbCount:    1,
			Reserved:    appNS.IsReserved(),
		}
	}

	addresses, err := createAddresses(addressesData{
		{t: storageTypes.MsgAddressTypeSigner, address: m.Signer},
	}, level)

	return storageTypes.MsgPayForBlobs, addresses, namespaces, blobsSize, err
}
