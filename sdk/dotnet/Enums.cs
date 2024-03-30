// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace UnMango.KubernetesTheHardWay
{
    /// <summary>
    /// CPU architecture
    /// </summary>
    [EnumType]
    public readonly struct Architecture : IEquatable<Architecture>
    {
        private readonly string _value;

        private Architecture(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Architecture Amd64 { get; } = new Architecture("amd64");
        public static Architecture Arm64 { get; } = new Architecture("arm64");

        public static bool operator ==(Architecture left, Architecture right) => left.Equals(right);
        public static bool operator !=(Architecture left, Architecture right) => !left.Equals(right);

        public static explicit operator string(Architecture value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Architecture other && Equals(other);
        public bool Equals(Architecture other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
