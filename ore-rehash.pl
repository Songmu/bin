#!/usr/bin/env perl
use 5.014;
use warnings;
use utf8;
use autodie;
use File::Spec;
use File::Basename;

my $bin = File::Spec->catdir($ENV{HOME}, 'bin');
chdir $bin;

opendir(my $dh, $bin) or die "can't opendir $bin: $!";
while (my $dir = readdir $dh) {
    next if ! -d $dir || $dir =~ /^\./;

    bulk_ln(File::Spec->catdir($bin, $dir));
}


sub bulk_ln {
    my $dir = shift;

    opendir(my $dh, $dir) or die "can't opendir $bin: $!";
    while (my $file = readdir $dh) {
        my $abs = File::Spec->catfile($dir, $file);
        next unless -f -x $abs;
        `ln -s -i $abs`;
    }
}

