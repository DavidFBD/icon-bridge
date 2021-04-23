// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2021 ICON Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity >=0.4.22 <0.8.5;

interface IBSH {
    /**
      @notice Handle BTP Message from other blockchain.
      @dev Accept the message only from the BMC.
      Every BSH must implement this function
      @param _from    Network Address of source network
      @param _svc     Name of the service
      @param _sn      Serial number of the message
      @param _msg     Serialized bytes of ServiceMessage
  */

    function handleBTPMessage(
        string calldata _from,
        string calldata _svc,
        uint256 _sn,
        bytes calldata _msg
    ) external;

    /**
      @notice Handle the error on delivering the message.
      @dev Accept the error only from the BMC.
      Every BSH must implement this function
      @param _src     BTP Address of BMC generates the error
      @param _svc     Name of the service
      @param _sn      Serial number of the original message
      @param _code    Code of the error
      @param _msg     Message of the error 
  */
    function handleBTPError(
        string calldata _src,
        string calldata _svc,
        uint256 _sn,
        uint256 _code,
        string calldata _msg
    ) external;
}
