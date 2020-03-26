// Copyright 2018-2020 Polyaxon, Inc.
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

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.6
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package io.swagger.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import io.swagger.client.model.V1ResourceRequirements;
import java.io.IOException;

/**
 * V1PolyaxonSidecarContainer
 */

public class V1PolyaxonSidecarContainer {
  @SerializedName("image")
  private String image = null;

  @SerializedName("image_tag")
  private String imageTag = null;

  @SerializedName("image_pull_policy")
  private String imagePullPolicy = null;

  @SerializedName("sleep_interval")
  private Integer sleepInterval = null;

  @SerializedName("sync_interval")
  private Integer syncInterval = null;

  @SerializedName("resources")
  private V1ResourceRequirements resources = null;

  public V1PolyaxonSidecarContainer image(String image) {
    this.image = image;
    return this;
  }

   /**
   * Get image
   * @return image
  **/
  @ApiModelProperty(value = "")
  public String getImage() {
    return image;
  }

  public void setImage(String image) {
    this.image = image;
  }

  public V1PolyaxonSidecarContainer imageTag(String imageTag) {
    this.imageTag = imageTag;
    return this;
  }

   /**
   * Get imageTag
   * @return imageTag
  **/
  @ApiModelProperty(value = "")
  public String getImageTag() {
    return imageTag;
  }

  public void setImageTag(String imageTag) {
    this.imageTag = imageTag;
  }

  public V1PolyaxonSidecarContainer imagePullPolicy(String imagePullPolicy) {
    this.imagePullPolicy = imagePullPolicy;
    return this;
  }

   /**
   * Get imagePullPolicy
   * @return imagePullPolicy
  **/
  @ApiModelProperty(value = "")
  public String getImagePullPolicy() {
    return imagePullPolicy;
  }

  public void setImagePullPolicy(String imagePullPolicy) {
    this.imagePullPolicy = imagePullPolicy;
  }

  public V1PolyaxonSidecarContainer sleepInterval(Integer sleepInterval) {
    this.sleepInterval = sleepInterval;
    return this;
  }

   /**
   * Get sleepInterval
   * @return sleepInterval
  **/
  @ApiModelProperty(value = "")
  public Integer getSleepInterval() {
    return sleepInterval;
  }

  public void setSleepInterval(Integer sleepInterval) {
    this.sleepInterval = sleepInterval;
  }

  public V1PolyaxonSidecarContainer syncInterval(Integer syncInterval) {
    this.syncInterval = syncInterval;
    return this;
  }

   /**
   * Get syncInterval
   * @return syncInterval
  **/
  @ApiModelProperty(value = "")
  public Integer getSyncInterval() {
    return syncInterval;
  }

  public void setSyncInterval(Integer syncInterval) {
    this.syncInterval = syncInterval;
  }

  public V1PolyaxonSidecarContainer resources(V1ResourceRequirements resources) {
    this.resources = resources;
    return this;
  }

   /**
   * Get resources
   * @return resources
  **/
  @ApiModelProperty(value = "")
  public V1ResourceRequirements getResources() {
    return resources;
  }

  public void setResources(V1ResourceRequirements resources) {
    this.resources = resources;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1PolyaxonSidecarContainer v1PolyaxonSidecarContainer = (V1PolyaxonSidecarContainer) o;
    return Objects.equals(this.image, v1PolyaxonSidecarContainer.image) &&
        Objects.equals(this.imageTag, v1PolyaxonSidecarContainer.imageTag) &&
        Objects.equals(this.imagePullPolicy, v1PolyaxonSidecarContainer.imagePullPolicy) &&
        Objects.equals(this.sleepInterval, v1PolyaxonSidecarContainer.sleepInterval) &&
        Objects.equals(this.syncInterval, v1PolyaxonSidecarContainer.syncInterval) &&
        Objects.equals(this.resources, v1PolyaxonSidecarContainer.resources);
  }

  @Override
  public int hashCode() {
    return Objects.hash(image, imageTag, imagePullPolicy, sleepInterval, syncInterval, resources);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1PolyaxonSidecarContainer {\n");
    
    sb.append("    image: ").append(toIndentedString(image)).append("\n");
    sb.append("    imageTag: ").append(toIndentedString(imageTag)).append("\n");
    sb.append("    imagePullPolicy: ").append(toIndentedString(imagePullPolicy)).append("\n");
    sb.append("    sleepInterval: ").append(toIndentedString(sleepInterval)).append("\n");
    sb.append("    syncInterval: ").append(toIndentedString(syncInterval)).append("\n");
    sb.append("    resources: ").append(toIndentedString(resources)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}
