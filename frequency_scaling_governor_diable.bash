#! /bin/bash
for i in /sys/devices/system/cpu/cpu[0-9]*
do
    echo "Disabling Frequency Scaling for CPU $i"
    echo performance > $i/cpufreq/scaling_governor
done
