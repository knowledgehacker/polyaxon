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
 *
 */

import ApiClient from '../ApiClient';

/**
 * The ProtobufAny model module.
 * @module model/ProtobufAny
 * @version 1.5.2
 */
class ProtobufAny {
    /**
     * Constructs a new <code>ProtobufAny</code>.
     * &#x60;Any&#x60; contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message.  Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type.  Example 1: Pack and unpack a message in C++.      Foo foo &#x3D; ...;     Any any;     any.PackFrom(foo);     ...     if (any.UnpackTo(&amp;foo)) {       ...     }  Example 2: Pack and unpack a message in Java.      Foo foo &#x3D; ...;     Any any &#x3D; Any.pack(foo);     ...     if (any.is(Foo.class)) {       foo &#x3D; any.unpack(Foo.class);     }   Example 3: Pack and unpack a message in Python.      foo &#x3D; Foo(...)     any &#x3D; Any()     any.Pack(foo)     ...     if any.Is(Foo.DESCRIPTOR):       any.Unpack(foo)       ...   Example 4: Pack and unpack a message in Go       foo :&#x3D; &amp;pb.Foo{...}      any, err :&#x3D; ptypes.MarshalAny(foo)      ...      foo :&#x3D; &amp;pb.Foo{}      if err :&#x3D; ptypes.UnmarshalAny(any, foo); err !&#x3D; nil {        ...      }  The pack methods provided by protobuf library will by default use &#39;type.googleapis.com/full.type.name&#39; as the type URL and the unpack methods only use the fully qualified type name after the last &#39;/&#39; in the type URL, for example \&quot;foo.bar.com/x/y.z\&quot; will yield type name \&quot;y.z\&quot;.   JSON &#x3D;&#x3D;&#x3D;&#x3D; The JSON representation of an &#x60;Any&#x60; value uses the regular representation of the deserialized, embedded message, with an additional field &#x60;@type&#x60; which contains the type URL. Example:      package google.profile;     message Person {       string first_name &#x3D; 1;       string last_name &#x3D; 2;     }      {       \&quot;@type\&quot;: \&quot;type.googleapis.com/google.profile.Person\&quot;,       \&quot;firstName\&quot;: &lt;string&gt;,       \&quot;lastName\&quot;: &lt;string&gt;     }  If the embedded message type is well-known and has a custom JSON representation, that representation will be embedded adding a field &#x60;value&#x60; which holds the custom JSON in addition to the &#x60;@type&#x60; field. Example (for message [google.protobuf.Duration][]):      {       \&quot;@type\&quot;: \&quot;type.googleapis.com/google.protobuf.Duration\&quot;,       \&quot;value\&quot;: \&quot;1.212s\&quot;     }
     * @alias module:model/ProtobufAny
     */
    constructor() { 
        
        ProtobufAny.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ProtobufAny</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ProtobufAny} obj Optional instance to populate.
     * @return {module:model/ProtobufAny} The populated <code>ProtobufAny</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ProtobufAny();

            if (data.hasOwnProperty('type_url')) {
                obj['type_url'] = ApiClient.convertToType(data['type_url'], 'String');
            }
            if (data.hasOwnProperty('value')) {
                obj['value'] = ApiClient.convertToType(data['value'], 'Blob');
            }
        }
        return obj;
    }


}

/**
 * A URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one \"/\" character. The last segment of the URL's path must represent the fully qualified name of the type (as in `path/google.protobuf.Duration`). The name should be in a canonical form (e.g., leading \".\" is not accepted).  In practice, teams usually precompile into the binary all types that they expect it to use in the context of Any. However, for URLs which use the scheme `http`, `https`, or no scheme, one can optionally set up a type server that maps type URLs to message definitions as follows:  * If no scheme is provided, `https` is assumed. * An HTTP GET on the URL must yield a [google.protobuf.Type][]   value in binary format, or produce an error. * Applications are allowed to cache lookup results based on the   URL, or have them precompiled into a binary to avoid any   lookup. Therefore, binary compatibility needs to be preserved   on changes to types. (Use versioned type names to manage   breaking changes.)  Note: this functionality is not currently available in the official protobuf release, and it is not used for type URLs beginning with type.googleapis.com.  Schemes other than `http`, `https` (or the empty scheme) might be used with implementation specific semantics.
 * @member {String} type_url
 */
ProtobufAny.prototype['type_url'] = undefined;

/**
 * Must be a valid serialized protocol buffer of the above specified type.
 * @member {Blob} value
 */
ProtobufAny.prototype['value'] = undefined;






export default ProtobufAny;

