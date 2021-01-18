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

/*
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


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import org.openapitools.client.model.V1CronSchedule;
import org.openapitools.client.model.V1DateTimeSchedule;
import org.openapitools.client.model.V1IntervalSchedule;

/**
 * V1Schedule
 */

public class V1Schedule {
  public static final String SERIALIZED_NAME_CRON = "cron";
  @SerializedName(SERIALIZED_NAME_CRON)
  private V1CronSchedule cron;

  public static final String SERIALIZED_NAME_DATETIME = "datetime";
  @SerializedName(SERIALIZED_NAME_DATETIME)
  private V1DateTimeSchedule datetime;

  public static final String SERIALIZED_NAME_INTERVAL = "interval";
  @SerializedName(SERIALIZED_NAME_INTERVAL)
  private V1IntervalSchedule interval;


  public V1Schedule cron(V1CronSchedule cron) {
    
    this.cron = cron;
    return this;
  }

   /**
   * Get cron
   * @return cron
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1CronSchedule getCron() {
    return cron;
  }


  public void setCron(V1CronSchedule cron) {
    this.cron = cron;
  }


  public V1Schedule datetime(V1DateTimeSchedule datetime) {
    
    this.datetime = datetime;
    return this;
  }

   /**
   * Get datetime
   * @return datetime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1DateTimeSchedule getDatetime() {
    return datetime;
  }


  public void setDatetime(V1DateTimeSchedule datetime) {
    this.datetime = datetime;
  }


  public V1Schedule interval(V1IntervalSchedule interval) {
    
    this.interval = interval;
    return this;
  }

   /**
   * Get interval
   * @return interval
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1IntervalSchedule getInterval() {
    return interval;
  }


  public void setInterval(V1IntervalSchedule interval) {
    this.interval = interval;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Schedule v1Schedule = (V1Schedule) o;
    return Objects.equals(this.cron, v1Schedule.cron) &&
        Objects.equals(this.datetime, v1Schedule.datetime) &&
        Objects.equals(this.interval, v1Schedule.interval);
  }

  @Override
  public int hashCode() {
    return Objects.hash(cron, datetime, interval);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Schedule {\n");
    sb.append("    cron: ").append(toIndentedString(cron)).append("\n");
    sb.append("    datetime: ").append(toIndentedString(datetime)).append("\n");
    sb.append("    interval: ").append(toIndentedString(interval)).append("\n");
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

