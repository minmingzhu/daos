/*
 * (C) Copyright 2018-2019 Intel Corporation.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * GOVERNMENT LICENSE RIGHTS-OPEN SOURCE SOFTWARE
 * The Government's rights to use, modify, reproduce, release, perform, display,
 * or disclose this software are subject to the terms of the Apache License as
 * provided in Contract No. B609815.
 * Any reproduction of computer software, computer software documentation, or
 * portions thereof marked with this legend must also reproduce the markings.
 */

package io.daos.obj;

import io.daos.DaosClient;
import io.daos.DaosObjectType;
import io.daos.ForceCloseable;
import io.daos.SharableClient;
import org.apache.commons.lang.ObjectUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.nio.ByteBuffer;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

/**
 * A sharable Java DAOS Object client to wrap all object related APIs.
 * {@link DaosObject} and {@link DaosObjectId} should be create from this client.
 * It registers itself to shutdown manager in {@link DaosClient} to release resources in case of abnormal shutdown.
 */
public class DaosObjClient extends SharableClient implements ForceCloseable {

  private long contPtr;

  // keyed by poolId+contId
  private static final Map<String, DaosObjClient> pcObjMap = new ConcurrentHashMap<>();

  private static final Logger log = LoggerFactory.getLogger(DaosObjClient.class);

  private DaosObjClient(String poolId, String contId, DaosObjClientBuilder builder) {
    super(poolId, contId, builder);
  }

  private synchronized void init() throws IOException {
    if (isInited()) {
      return;
    }
    DaosObjClientBuilder builder = getBuilder();
    setClient(builder.buildDaosClient());
    DaosClient client = getClient();
    contPtr = client.getContPtr();
    client.registerForShutdown(this);
    setInited(true);
    log.info("DaosObjClient for {}, {} initialized", builder.getPoolId(), builder.getContId());
  }

  /**
   * create new instance of un-encoded {@link DaosObjectId} with high and low value of 0.
   *
   * @return DaosObjectId object
   */
  public DaosObjectId newObjectId() {
    return new DaosObjectId(contPtr);
  }

  /**
   * create new instance of encoded {@link DaosObjectId} with high and low value of 0.
   *
   * @param feats
   * object feature fits
   * @param objectType
   * object type
   * @param args
   * reserved
   * @return DaosObjectId object
   */
  public DaosObjectId newEncodedObjectId(int feats, DaosObjectType objectType, int args) {
    DaosObjectId id = new DaosObjectId(contPtr);
    id.encode(feats, objectType, args);
    return id;
  }

  /**
   * create new instance of un-encoded {@link DaosObjectId} with specified high and low values.
   *
   * @param high
   * high value of ID
   * @param low
   * low value of ID
   * @return DaosObjectId object
   */
  public DaosObjectId newObjectId(long high, long low) {
    return new DaosObjectId(contPtr, high, low);
  }

  /**
   * create new instance of encoded {@link DaosObjectId} with specified high and low values.
   *
   * @param high
   * high value of ID
   * @param low
   * low value of ID
   * @param feats
   * object feature bits
   * @param objectType
   * object type
   * @param args
   * reserved
   * @return DaosObjectId object
   */
  public DaosObjectId newEncodedObjectId(long high, long low, int feats, DaosObjectType objectType, int args) {
    DaosObjectId id = new DaosObjectId(contPtr, high, low);
    id.encode(feats, objectType, args);
    return id;
  }

  /**
   * encode object id with object feature bits and object type.
   * encoded object id is set back to <code>buffer</code>.
   *
   * @param contPtr
   * opened container handler
   * @param oidBuffer
   * direct byte buffer with original object id's high and low. encode object id is set back to this buffer.
   * @param feats
   * object feature bits
   * @param objectTypeName
   * object type name, see {@link DaosObjectType}
   * @param args
   * reserved
   */
  public native static void encodeObjectId(long contPtr, ByteBuffer oidBuffer, int feats, String objectTypeName,
                                           int args);

  /**
   * open object denoted by object id stored in <code>buffer</code>.
   *
   * @param contPtr
   * opened container handler
   * @param oidBuffer
   * direct byte buffer with original object id's high and low
   * @param mode
   * open mode, see {@link OpenMode}
   * @return handle of opened object
   */
  public native static long openObject(long contPtr, ByteBuffer oidBuffer, int mode);

  /**
   * close object.
   *
   * @param objectPtr
   * handle of opened object
   */
  public native static void closeObject(long objectPtr);

  /**
   * punch an entire object with all associated with it.
   *
   * @param objectPtr
   * handle of opened object
   * @param flags
   * punch flags (currently ignored)
   */
  public native static void punchObject(long objectPtr, long flags);

  /**
   * punch dkeys (with all its akeys) from an object.
   *
   * @param objectPtr
   * handle of opened object
   * @param flags
   * punch flags (currently ignored)
   * @param nbrOfDkeys
   * number of dkeys
   * @param dkeysBuffer
   * dkeys written into direct byte buffer in format, len1+key1+len2+key2...
   */
  public native static void punchObjectDkeys(long objectPtr, long flags, int nbrOfDkeys, ByteBuffer dkeysBuffer);

  /**
   * punch akeys (with all records) from an object.
   * 
   * @param objectPtr
   * handle of opened object
   * @param flags
   * punch flags (currently ignored)
   * @param nbrOfAkeys
   * number of akeys
   * @param keysBuffer
   * dkey and akeys written into direct byte buffer in format, dkey len+dkey+akey1 len+akey1+akey2 len+akey2...
   */
  public native static void punchObjectAkeys(long objectPtr, long flags, int nbrOfAkeys, ByteBuffer keysBuffer);

  /**
   * query attributes of an object.
   *
   * @param objectPtr
   * handle of opened object
   * @return attributes serialized by protobuf. see DaosObjectAttribute.proto.
   */
  public native static byte[] queryObjectAttribute(long objectPtr);

  @Override
  protected synchronized void disconnect(boolean force) throws IOException {
    decrementRef();
    DaosObjClientBuilder builder = getBuilder();
    if ((force || getRefCnt() <= 0) && isInited()) {
      if (force) {
        getClient().forceClose();
      } else {
        getClient().close();
      }
      log.info("DaosObjClient for {}, {} disconnected", builder.getPoolId(), builder.getContId());
    }
    setInited(false);
    pcObjMap.remove(builder.getPoolId() + builder.getContId());
  }

  @Override
  protected DaosObjClientBuilder getBuilder() {
    return (DaosObjClientBuilder)super.getBuilder();
  }

  public static class DaosObjClientBuilder extends DaosClient.DaosClientBuilder<DaosObjClientBuilder> {
    private boolean shareFsClient;

    /**
     * share {@link DaosObjClient} instance or not.
     *
     * @param shareFsClient
     * default is true
     * @return DaosFsClientBuilder
     */
    public DaosObjClientBuilder shareFsClient(boolean shareFsClient) {
      this.shareFsClient = shareFsClient;
      return this;
    }

    @Override
    public DaosObjClientBuilder clone() throws CloneNotSupportedException {
      return (DaosObjClientBuilder) super.clone();
    }

    @Override
    public DaosObjClient build() throws IOException {
      String poolId = getPoolId();
      String contId = getContId();
      DaosObjClientBuilder builder = (DaosObjClientBuilder) ObjectUtils.clone(this);
      DaosObjClient objClient;
      if (!builder.shareFsClient) {
        objClient = new DaosObjClient(poolId, contId, builder);
      } else {
        //check existing client
        if (poolId == null) {
          throw new IllegalArgumentException("need pool UUID.");
        }
        if (contId == null) {
          throw new IllegalArgumentException("need container UUID.");
        }
        String key = poolId + contId;
        objClient = pcObjMap.get(key);
        if (objClient == null) {
          objClient = new DaosObjClient(poolId, contId, builder);
          pcObjMap.putIfAbsent(key, objClient);
        }
        objClient = pcObjMap.get(key);
      }
      objClient.init();
      objClient.incrementRef();
      return objClient;
    }

    protected DaosClient buildDaosClient() throws IOException {
      return (DaosClient) super.build();
    }
  }
}