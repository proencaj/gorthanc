package gorthanc

import (
	"fmt"

	"github.com/proencaj/gorthanc/types"
)

func (c *Client) GetPeers() ([]string, error) {
	var peers []string
	if err := c.get("peers", &peers); err != nil {
		return nil, err
	}
	return peers, nil
}

func (c *Client) GetPeerDetails(peerName string) (*types.Peer, error) {
	var peer types.Peer
	path := fmt.Sprintf("peers/%s/configuration", peerName)

	if err := c.get(path, &peer); err != nil {
		return nil, err
	}

	return &peer, nil
}

func (c *Client) CreateOrUpdatePeer(peerName string, request *types.PeerCreateRequest) error {
	path := fmt.Sprintf("peers/%s", peerName)

	if err := c.put(path, request, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) DeletePeer(peerName string) error {
	path := fmt.Sprintf("peers/%s", peerName)

	if err := c.delete(path, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) StoreToPeer(peerName, resourceID string) error {
	path := fmt.Sprintf("peers/%s/store", peerName)

	if err := c.post(path, resourceID, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) StoreToPeerWithOptions(peerName string, request *types.PeerStoreRequest) (*types.PeerStoreResult, error) {
	path := fmt.Sprintf("peers/%s/store", peerName)

	var result types.PeerStoreResult
	if err := c.post(path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPeerSystem(peerName string) (*types.SystemInfo, error) {
	var info types.SystemInfo
	path := fmt.Sprintf("peers/%s/system", peerName)

	if err := c.get(path, &info); err != nil {
		return nil, err
	}

	return &info, nil
}
