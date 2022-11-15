# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: wallet.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0cwallet.proto\x12\x06pactus\"}\n\x13\x43reateWalletRequest\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n\x08mnemonic\x18\x02 \x01(\tR\x08mnemonic\x12\x1a\n\x08language\x18\x03 \x01(\tR\x08language\x12\x1a\n\x08password\x18\x04 \x01(\tR\x08password\"\x16\n\x14\x43reateWalletResponse\"\'\n\x11LoadWalletRequest\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\"(\n\x12LoadWalletResponse\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\")\n\x13UnloadWalletRequest\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\"*\n\x14UnloadWalletResponse\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\"I\n\x11LockWalletRequest\x12\x1a\n\x08password\x18\x01 \x01(\tR\x08password\x12\x18\n\x07timeout\x18\x02 \x01(\x05R\x07timeout\"\x14\n\x12LockWalletResponse\"K\n\x13UnlockWalletRequest\x12\x1a\n\x08password\x18\x01 \x01(\tR\x08password\x12\x18\n\x07timeout\x18\x02 \x01(\x05R\x07timeout\"\x16\n\x14UnlockWalletResponse2\xf3\x02\n\x06Wallet\x12I\n\x0c\x43reateWallet\x12\x1b.pactus.CreateWalletRequest\x1a\x1c.pactus.CreateWalletResponse\x12\x43\n\nLoadWallet\x12\x19.pactus.LoadWalletRequest\x1a\x1a.pactus.LoadWalletResponse\x12I\n\x0cUnloadWallet\x12\x1b.pactus.UnloadWalletRequest\x1a\x1c.pactus.UnloadWalletResponse\x12\x43\n\nLockWallet\x12\x19.pactus.LockWalletRequest\x1a\x1a.pactus.LockWalletResponse\x12I\n\x0cUnlockWallet\x12\x1b.pactus.UnlockWalletRequest\x1a\x1c.pactus.UnlockWalletResponseBA\n\rpactus.walletZ0github.com/pactus-project/pactus/www/grpc/pactusb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'wallet_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\rpactus.walletZ0github.com/pactus-project/pactus/www/grpc/pactus'
  _CREATEWALLETREQUEST._serialized_start=24
  _CREATEWALLETREQUEST._serialized_end=149
  _CREATEWALLETRESPONSE._serialized_start=151
  _CREATEWALLETRESPONSE._serialized_end=173
  _LOADWALLETREQUEST._serialized_start=175
  _LOADWALLETREQUEST._serialized_end=214
  _LOADWALLETRESPONSE._serialized_start=216
  _LOADWALLETRESPONSE._serialized_end=256
  _UNLOADWALLETREQUEST._serialized_start=258
  _UNLOADWALLETREQUEST._serialized_end=299
  _UNLOADWALLETRESPONSE._serialized_start=301
  _UNLOADWALLETRESPONSE._serialized_end=343
  _LOCKWALLETREQUEST._serialized_start=345
  _LOCKWALLETREQUEST._serialized_end=418
  _LOCKWALLETRESPONSE._serialized_start=420
  _LOCKWALLETRESPONSE._serialized_end=440
  _UNLOCKWALLETREQUEST._serialized_start=442
  _UNLOCKWALLETREQUEST._serialized_end=517
  _UNLOCKWALLETRESPONSE._serialized_start=519
  _UNLOCKWALLETRESPONSE._serialized_end=541
  _WALLET._serialized_start=544
  _WALLET._serialized_end=915
# @@protoc_insertion_point(module_scope)
