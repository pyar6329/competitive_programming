{-# LANGUAGE DuplicateRecordFields, FlexibleInstances, UndecidableInstances #-}

module Main where

import Control.Monad
import Data.Array
import Data.Bits
import Data.List
import Data.List.Split
import Data.Set
import Data.Text
import Debug.Trace
import System.Environment
import System.IO
import System.IO.Unsafe

-- Complete the compareTriplets function below.
compareTriplets :: [Int] -> [Int] -> [Int]
compareTriplets a b = do
    let pairRank = Data.List.zip a b
    let aliceWins = [i > j | (i, j) <- pairRank ]
    let alice = Data.List.length $ Data.List.filter (\x -> x == True) aliceWins
    let bobWins = [i < j | (i, j) <- pairRank ]
    let bob = Data.List.length $ Data.List.filter (\x -> x == True) bobWins
    [alice, bob]

lstrip = Data.Text.unpack . Data.Text.stripStart . Data.Text.pack
rstrip = Data.Text.unpack . Data.Text.stripEnd . Data.Text.pack

main :: IO()
main = do
    stdout <- getEnv "OUTPUT_PATH"
    fptr <- openFile stdout WriteMode

    aTemp <- getLine

    let a = Data.List.map (read :: String -> Int) . Data.List.words $ rstrip aTemp

    bTemp <- getLine

    let b = Data.List.map (read :: String -> Int) . Data.List.words $ rstrip bTemp

    let result = compareTriplets a b

    hPutStrLn fptr $ Data.List.intercalate " " $ Data.List.map (\x -> show x) $ result

    hFlush fptr
    hClose fptr
