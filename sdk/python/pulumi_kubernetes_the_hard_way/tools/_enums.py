# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'SystemctlCommand',
    'TeeMode',
]


class SystemctlCommand(str, Enum):
    BIND = "bind"
    CAT = "cat"
    CLEAN = "clean"
    DAEMON_RELOAD = "daemon-reload"
    DISABLE = "disable"
    ENABLE = "enable"
    FREEZE = "freeze"
    IS_ACTIVE = "is-active"
    IS_ENABLED = "is-enabled"
    IS_FAILED = "is-failed"
    ISOLATE = "isolate"
    KILL = "kill"
    LIST_AUTOMOUNTS = "list-automounts"
    LIST_DEPENDENCIES = "list-dependencies"
    LIST_PATHS = "list-paths"
    LIST_SOCKETS = "list-sockets"
    LIST_TIMERS = "list-timers"
    LIST_UNITS = "list-units"
    MASK = "mask"
    MOUNT_IMAGE = "mount-image"
    REENABLE = "reenable"
    RELOAD = "reload"
    RELOAD_OR_RESTART = "reload-or-restart"
    RESTART = "restart"
    SET_PROPERTY = "set-property"
    SHOW = "show"
    START = "start"
    STATUS = "status"
    STOP = "stop"
    THAW = "thaw"
    TRY_RELOAD_OR_RESTART = "try-reload-or-restart"
    TRY_RESTART = "try-restart"
    UNMASK = "unmask"


class TeeMode(str, Enum):
    WARN = "warn"
    WARN_NOPIPE = "warn-nopipe"
    EXIT = "exit"
    EXIT_NOPIPE = "exit-nopipe"
