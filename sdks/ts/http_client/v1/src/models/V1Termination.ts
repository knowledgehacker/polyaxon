// Copyright 2018-2021 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.5.2
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * Termination specification, it defines how to handle failures, job runtime, and cleanup policy.
 * @export
 * @interface V1Termination
 */
export interface V1Termination {
    /**
     * 
     * @type {number}
     * @memberof V1Termination
     */
    max_retries?: number;
    /**
     * 
     * @type {number}
     * @memberof V1Termination
     */
    ttl?: number;
    /**
     * 
     * @type {number}
     * @memberof V1Termination
     */
    timeout?: number;
}

export function V1TerminationFromJSON(json: any): V1Termination {
    return V1TerminationFromJSONTyped(json, false);
}

export function V1TerminationFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Termination {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'max_retries': !exists(json, 'max_retries') ? undefined : json['max_retries'],
        'ttl': !exists(json, 'ttl') ? undefined : json['ttl'],
        'timeout': !exists(json, 'timeout') ? undefined : json['timeout'],
    };
}

export function V1TerminationToJSON(value?: V1Termination | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'max_retries': value.max_retries,
        'ttl': value.ttl,
        'timeout': value.timeout,
    };
}


